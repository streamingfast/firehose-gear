package rpc

import (
	"fmt"
	"slices"

	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/types"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

type MetadataConverter struct {
	gearClients *firecoreRPC.Clients[*Client]
	logger      *zap.Logger
	tracer      logging.Tracer
}

func NewMetadataConverter(gearClients *firecoreRPC.Clients[*Client], logger *zap.Logger, tracer logging.Tracer) *MetadataConverter {
	return &MetadataConverter{
		gearClients: gearClients,
		logger:      logger,
		tracer:      tracer,
	}
}

func (mc *MetadataConverter) Convert(blockHash string) (*types.Types, error) {
	metadata, err := mc.fetchStateMetadata(blockHash)
	if err != nil {
		return nil, fmt.Errorf("fetching state metadata: %w", err)
	}

	if metadata.Version != 14 {
		return nil, nil
	}

	switch metadata.Version {
	case 14:
		return convertTypesFromv14(metadata.AsMetadataV14)
	default:
		fmt.Println("Unsupported metadata version", metadata.Version)
	}

	return nil, nil
}

func (mc *MetadataConverter) fetchStateMetadata(blockHash string) (*substrateTypes.Metadata, error) {
	return firecoreRPC.WithClients(mc.gearClients, func(client *Client) (*substrateTypes.Metadata, error) {
		if blockHash == "" {
			return client.state.GetMetadataLatest()
		}

		bHash, err := substrateTypes.NewHashFromHexString(blockHash)
		if err != nil {
			return nil, fmt.Errorf("parsing block hash: %w", err)
		}

		return client.state.GetMetadata(bHash)
	})
}

func convertTypesFromv14(metadata substrateTypes.MetadataV14) (*types.Types, error) {
	t := &types.Types{}
	ttypes := make([]types.IType, 0)
	allMetadataTypes := metadata.Lookup.Types
	pallets := metadata.Pallets

	// Utility, Whitelist, Scheduler, Proxy, Multisig
	// Any pallet that contains the call type, will cause and infinite loop
	skip := []int{15, 21, 22, 25, 26}

	for i, pallet := range pallets {
		if slices.Contains(skip, i) {
			continue
		}

		if pallet.HasCalls { // CallFunctions
			call := pallet.Calls // this is to fetch in the metadata types to get the CallFunctions
			id := call.Type.Int64()
			portableType := allMetadataTypes[id]
			convertType := &ConvertType{}
			ttype := convertType.convertPortableTypeV14(portableType, allMetadataTypes)
			variants := ttype.(*types.Variants)
			callFunctions := variants.ToCallFunctions()
			str := stringy.New(string(pallet.Name))
			protobufMessage := &types.ProtobufMessage{
				Name:          str.PascalCase().Get(),
				Messages:      variants.ToProtobufMessages(),
				CallFunctions: callFunctions,
			}
			ttypes = append(ttypes, protobufMessage)
		}
	}

	t.Types = ttypes
	return t, nil
}

/*
	- loop the Pallets
		- each Name will be the message name in protobuf
		- loop the items -> storage functions // todo: maybe we do not care about this
		- loop the calls -> call functions
		- loop Events -> events
		- loop Constants -> constants
		- loop Errors -> errors

{
  "Name": "Timestamp.set",
  "CallFields": [
    {
      "Name": "now",
      "Value": 1719320625000,
      "LookupIndex": 10 <- find this in the types.[10]
    }
  ],
  "CallIndex": {
    "SectionIndex": 1, <- is the inddex in the pallets
    "MethodIndex": 0 <- is the index in type definition, for example variants.[0]
  },
  "Version": 4,
  "Signature": null
}

*/

type ConvertType struct {
	name string
}

func (ct *ConvertType) UpdateName(name string) {
	ct.name = name
}

func (ct *ConvertType) convertPortableTypeV14(ttype substrateTypes.PortableTypeV14, allMetadataTypes []substrateTypes.PortableTypeV14) types.IType {
	if ttype.Type.Def.IsComposite {
		c := ttype.Type.Def.Composite
		return ct.convertCompositeType(c, allMetadataTypes)
	}

	if ttype.Type.Def.IsVariant {
		v := ttype.Type.Def.Variant
		return ct.convertVariantType(v, allMetadataTypes)
	}

	if ttype.Type.Def.IsSequence {
		s := ttype.Type.Def.Sequence
		return ct.convertSequenceType(s, allMetadataTypes)
	}

	if ttype.Type.Def.IsArray {
		a := ttype.Type.Def.Array
		return ct.convertArrayType(a, allMetadataTypes)
	}

	if ttype.Type.Def.IsTuple {
		t := ttype.Type.Def.Tuple
		return ct.convertTupleType(t, allMetadataTypes)
	}

	if ttype.Type.Def.IsPrimitive {
		p := ttype.Type.Def.Primitive
		b := substrateTypes.Si0TypeDefPrimitive(p.Si0TypeDefPrimitive)
		return ct.convertPrimitiveType(b)
	}

	if ttype.Type.Def.IsCompact {
		c := ttype.Type.Def.Compact
		return ct.convertCompactType(c, allMetadataTypes)
	}

	if ttype.Type.Def.IsBitSequence {
		b := ttype.Type.Def.BitSequence
		return ct.convertBitSequenceType(b, allMetadataTypes)
	}

	if ttype.Type.Def.IsHistoricMetaCompat {
		h := ttype.Type.Def.HistoricMetaCompat
		return ct.convertHistoricMetaCompat(h)
	}

	panic(fmt.Sprintf("unsupported type: %v", ttype.Type.Def))
}

func (ct *ConvertType) convertCompositeType(c substrateTypes.Si1TypeDefComposite, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Composite {
	comp := &types.Composite{}
	fields := make([]*types.CompositeField, 0)

	for _, field := range c.Fields {
		lookupId := field.Type.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)

		fields = append(fields, &types.CompositeField{
			Name: string(field.Name),
			Type: convertedType,
		})
	}

	comp.CompositeFields = fields
	return comp
}

// NOTE: inside a variant we can find the definition of a Rust Result type, which would need to be handled a bit funkier...
func (ct *ConvertType) convertVariantType(v substrateTypes.Si1TypeDefVariant, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Variants {
	variant := &types.Variants{}
	variants := make([]*types.Variant, 0)

	for _, v := range v.Variants {
		fields := make([]*types.VariantField, 0)
		for _, field := range v.Fields {
			lookupId := field.Type.Int64()
			portableType := allMetadataTypes[lookupId]
			name := string(field.Name)
			ct.UpdateName(name)
			convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)

			fields = append(fields, &types.VariantField{
				Name: name,
				Type: convertedType,
			})
		}

		variants = append(variants, &types.Variant{
			Name:          string(v.Name),
			VariantFields: fields,
			Index:         uint64(v.Index) + 1, // this is the index in the protobuf, todo: do we still need this?
		})
	}

	variant.Variants = variants
	return variant
}

func (ct *ConvertType) convertSequenceType(s substrateTypes.Si1TypeDefSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Sequence {
	seq := &types.Sequence{}
	lookupId := s.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)
	seq.Item = convertedType
	return seq
}

func (ct *ConvertType) convertArrayType(a substrateTypes.Si1TypeDefArray, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Array {
	arr := &types.Array{}
	lookupId := a.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)
	arr.Type = convertedType
	return arr
}

func (ct *ConvertType) convertTupleType(t substrateTypes.Si1TypeDefTuple, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Tuple {
	tup := &types.Tuple{
		Name: ct.name,
	}
	items := make([]types.IType, 0)

	for _, item := range t {
		lookupId := item.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)
		items = append(items, convertedType)
	}

	tup.Items = items
	return tup
}

func (ct *ConvertType) convertPrimitiveType(b substrateTypes.Si0TypeDefPrimitive) *types.Primitive {
	p := &types.Primitive{}

	switch b {
	case substrateTypes.IsBool:
		p.Si0TypeDefPrimitive = &types.Type{IsBool: true}
	case substrateTypes.IsU8:
		p.Si0TypeDefPrimitive = &types.Type{IsU8: true}
	case substrateTypes.IsU16:
		p.Si0TypeDefPrimitive = &types.Type{IsU16: true}
	case substrateTypes.IsU32:
		p.Si0TypeDefPrimitive = &types.Type{IsU32: true}
	case substrateTypes.IsU64:
		p.Si0TypeDefPrimitive = &types.Type{IsU64: true}
	case substrateTypes.IsI8:
		p.Si0TypeDefPrimitive = &types.Type{IsI8: true}
	case substrateTypes.IsI16:
		p.Si0TypeDefPrimitive = &types.Type{IsI16: true}
	case substrateTypes.IsI32:
		p.Si0TypeDefPrimitive = &types.Type{IsI32: true}
	case substrateTypes.IsI64:
		p.Si0TypeDefPrimitive = &types.Type{IsI64: true}
	case substrateTypes.IsStr:
		p.Si0TypeDefPrimitive = &types.Type{IsStr: true}
	case substrateTypes.IsChar:
		p.Si0TypeDefPrimitive = &types.Type{IsChar: true}
	case substrateTypes.IsU128:
		p.Si0TypeDefPrimitive = &types.Type{IsU128: true}
	case substrateTypes.IsU256:
		p.Si0TypeDefPrimitive = &types.Type{IsU256: true}
	case substrateTypes.IsI128:
		p.Si0TypeDefPrimitive = &types.Type{IsI128: true}
	case substrateTypes.IsI256:
		p.Si0TypeDefPrimitive = &types.Type{IsI256: true}
	default:
		panic(fmt.Sprintf("unsupported primitive type %v", b))
	}

	return p
}

func (ct *ConvertType) convertCompactType(c substrateTypes.Si1TypeDefCompact, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Compact {
	comp := &types.Compact{}
	lookupId := c.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := ct.convertPortableTypeV14(portableType, allMetadataTypes)
	comp.Type = convertedType
	return comp
}

func (ct *ConvertType) convertBitSequenceType(b substrateTypes.Si1TypeDefBitSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.BitSequence {
	bitSeq := &types.BitSequence{}
	bitOrderPortableType := allMetadataTypes[b.BitOrderType.Int64()]
	bitStorePortableType := allMetadataTypes[b.BitStoreType.Int64()]
	bitSeq.BitOrderType = ct.convertPortableTypeV14(bitOrderPortableType, allMetadataTypes)
	bitSeq.BitStoreType = ct.convertPortableTypeV14(bitStorePortableType, allMetadataTypes)
	return bitSeq
}

func (ct *ConvertType) convertHistoricMetaCompat(h substrateTypes.Type) types.HistoricMetaCompat {
	return types.HistoricMetaCompat(string(h))
}
