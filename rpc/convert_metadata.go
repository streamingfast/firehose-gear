package rpc

import (
	"fmt"
	"strings"

	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/protobuf"
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

func (mc *MetadataConverter) Convert(blockHash string) (map[string]types.IType, error) {
	metadata, err := mc.fetchStateMetadata(blockHash)
	if err != nil {
		return nil, fmt.Errorf("fetching state metadata: %w", err)
	}

	if metadata.Version != 14 {
		return nil, nil
	}

	switch metadata.Version {
	case 14:
		tc := &TypeConverter{
			types:        make(map[string]types.IType),
			seenTypes:    make(map[int64]types.IType),
			seenMessages: make(map[string]*protobuf.Message),
		}
		return tc.convertTypesFromv14(metadata.AsMetadataV14)
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

type TypeConverter struct {
	// here we pass in the id of the types map, if we have already SEEN the id, we skip
	// nothing to do, we have already processed it
	types                              map[string]types.IType
	seenTypes                          map[int64]types.IType
	seenMessages                       map[string]*protobuf.Message
	allMetadataTypes                   []substrateTypes.PortableTypeV14
	currentFieldNameToBeRemovedSomehow string
}

func (tc *TypeConverter) UpdateName(name string) {
	// tc.name = name
}

func (c *TypeConverter) convertTypesFromv14(metadata substrateTypes.MetadataV14) (map[string]types.IType, error) {
	allMetadataTypes := metadata.Lookup.Types
	c.allMetadataTypes = allMetadataTypes

	for i, pallet := range metadata.Pallets {
		if i > 10000 {
			break
		}
		// Call
		if pallet.HasCalls {
			idx := pallet.Calls.Type.Int64()

			variants := c.allMetadataTypes[idx]

			calls := variants.Type.Def.Variant
			for _, variant := range calls.Variants {
				str := stringy.New(string(variant.Name) + "_Call")
				message := &protobuf.Message{
					Name: str.PascalCase().Get(), // identity.set_subs
				}
				c.ProcessCallFields(variant, message)
				c.seenMessages[message.Name] = message
			}
		}
	}

	for _, seenMsg := range c.seenMessages {
		fmt.Println(seenMsg.ToProto())
	}

	return nil, nil
}

func (c *TypeConverter) ProcessCallFields(variant substrateTypes.Si1Variant, message *protobuf.Message) {
	for _, f := range variant.Fields {
		fieldName := string(f.Name) // identity.set_subs.subs
		c.currentFieldNameToBeRemovedSomehow = fieldName
		field := c.ProcessField(f) // this should return the interface Field (either basic or optional)
		if field != nil {
			field.SetName(fieldName)
			message.Fields = append(message.Fields, field)
		}
		c.currentFieldNameToBeRemovedSomehow = "WARK"
	}
}

func (c *TypeConverter) ProcessField(f substrateTypes.Si1Field) protobuf.Field {
	idx := f.Type.Int64()
	ttype := c.allMetadataTypes[idx]

	return c.ProcessFieldType(ttype) // 177
}

func (c *TypeConverter) ProcessFieldType(ttype substrateTypes.PortableTypeV14) protobuf.Field {
	if len(ttype.Type.Path) != 0 {
		isOptional := ttype.Type.Path[len(ttype.Type.Path)-1]
		if isOptional == "Option" {
			ttype := c.ProcessOptionalType(ttype)
			return &protobuf.BasicField{
				Optional: true,
				Type:     ttype,
			}
		}
	}

	if ttype.Type.Def.IsPrimitive {
		return &protobuf.BasicField{
			Type: c.convertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).GetProtoFieldName(),
		}
	}

	if ttype.Type.Def.IsSequence {
		lookupId := ttype.Type.Def.Sequence.Type.Int64()
		lookupType := c.allMetadataTypes[lookupId] // 178
		typeName := c.ExtractTypeName(lookupType)

		return &protobuf.RepeatedField{
			Type: typeName,
		}
	}

	if ttype.Type.Def.IsTuple {
		return c.ProcessTupleType(ttype.Type.Def.Tuple)
	}

	if ttype.Type.Def.IsVariant {
		return c.ProcessVariantType(ttype.Type.Def.Variant)
	}

	// panic(fmt.Sprintf("unsupported type: %v", ttype.Type.Def))
	return nil
}

func (c *TypeConverter) ExtractTypeName(ttype substrateTypes.PortableTypeV14) string {
	name := ""
	if ttype.Type.Def.IsPrimitive {
		name = c.convertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).GetProtoFieldName()
	}

	if ttype.Type.Def.IsTuple {
		name = c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple)
	}

	if ttype.Type.Def.IsComposite {
		name = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsVariant {
		name = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsSequence {
		name = fmt.Sprintf("%s_list", c.currentFieldNameToBeRemovedSomehow)
	}

	if _, ok := c.seenMessages[name]; ok {
		return name
	}

	return name
}

func (c *TypeConverter) ExtractTypeNameFromTuple(tuple substrateTypes.Si1TypeDefTuple) string {
	var name strings.Builder
	name.WriteString("Tuple_")
	for _, item := range tuple {
		lookupId := item.Int64()
		ttype := c.allMetadataTypes[lookupId]
		name.WriteString(c.ExtractTypeName(ttype))
	}

	return name.String()
}

func (c *TypeConverter) ExtractTypeNameFromPath(path substrateTypes.Si1Path) string {
	parts := make([]string, 0)
	for _, p := range path {
		parts = append(parts, string(p))
	}
	return strings.Join(parts, "_")
}

func (c *TypeConverter) ProcessTupleType(tuple substrateTypes.Si1TypeDefTuple) *protobuf.BasicField {
	field := &protobuf.BasicField{
		Type: c.ExtractTypeNameFromTuple(tuple),
	}

	return field
}

func (c *TypeConverter) ProcessVariantType(variant substrateTypes.Si1TypeDefVariant) *protobuf.OneOfField {
	field := &protobuf.OneOfField{}
	for _, v := range variant.Variants {
		for _, f := range v.Fields {
			ttype := c.allMetadataTypes[f.Type.Int64()]

			if len(ttype.Type.Path) != 0 {
				typeName := c.ExtractTypeNameFromPath(ttype.Type.Path)
				field.Types = append(field.Types, typeName)
				continue
			}

			field.Types = append(field.Types, c.ExtractTypeName(ttype))
		}
	}

	return field
}

func (c *TypeConverter) ProcessOptionalType(ttype substrateTypes.PortableTypeV14) string {
	variant := ttype.Type.Def.Variant
	some := variant.Variants[1]
	someField := some.Fields[0]
	tttype := c.allMetadataTypes[someField.Type.Int64()]

	return c.ExtractTypeName(tttype)
}

func (c *TypeConverter) GetTypeNames(ttype types.IType) []string {
	switch v := ttype.(type) {
	case *types.Variants:
		variantNames := make([]string, 0)
		baseVariantName := v.GetName() // Variants

		for _, variant := range v.Variants {
			nestedVariantName := v.GetName() // Variant
			variantFieldTypeNames := make([]string, 0)

			for _, field := range variant.VariantFields {
				variantFieldName := field.GetName() // VariantField

				typeNames := c.GetTypeNames(field.Type)
				for _, typeName := range typeNames {
					variantFieldTypeNames = append(variantFieldTypeNames, fmt.Sprintf("%s%s", variantFieldName, typeName))
				}
			}

			for _, typeName := range variantFieldTypeNames {
				variantNames = append(variantNames, fmt.Sprintf("%s%s", nestedVariantName, typeName))
			}
		}

		for i, typeName := range variantNames {
			variantNames[i] = fmt.Sprintf("%s%s", baseVariantName, typeName)
		}

		return variantNames

	case *types.Composite:
		compositeTypeNames := make([]string, 0)
		baseCompositeName := v.GetName() // Composite

		for _, field := range v.CompositeFields {
			compositeFieldName := field.GetName() // CompositeField

			typeNames := c.GetTypeNames(field.Type)
			for _, typeName := range typeNames {
				compositeTypeNames = append(compositeTypeNames, fmt.Sprintf("%s%s", compositeFieldName, typeName))
			}
		}

		for i, typeName := range compositeTypeNames {
			compositeTypeNames[i] = fmt.Sprintf("%s%s", baseCompositeName, typeName)
		}

		return compositeTypeNames

	case *types.Tuple:
		baseName := v.GetName()
		tupeNames := make([]string, 0)

		for _, item := range v.Items {
			typeNames := c.GetTypeNames(item)

			for _, typeName := range typeNames {
				tupeNames = append(tupeNames, fmt.Sprintf("%s%s", baseName, typeName))
			}
		}

		return tupeNames

	case *types.Compact:
		baseName := v.GetName()
		compactNames := make([]string, 0)

		typeNames := c.GetTypeNames(v.Type)
		for _, typeName := range typeNames {
			compactNames = append(compactNames, fmt.Sprintf("%s%s", baseName, typeName))
		}

		return compactNames

	case *types.Array:
		baseName := v.GetName()
		arrayNames := make([]string, 0)

		typeNames := c.GetTypeNames(v.Type)
		for _, typeName := range typeNames {
			arrayNames = append(arrayNames, fmt.Sprintf("%s%s", baseName, typeName))
		}

		return arrayNames

	case *types.Sequence:
		baseName := v.GetName()
		sequenceNames := make([]string, 0)

		typeNames := c.GetTypeNames(v.Item)
		for _, typeName := range typeNames {
			sequenceNames = append(sequenceNames, fmt.Sprintf("%s%s", baseName, typeName))
		}

		return sequenceNames

	case *types.BitSequence:
		return []string{v.GetName()}

	case *types.HistoricMetaCompat:
		return []string{v.GetName()}

	case *types.Primitive:
		return []string{v.GetName()}

	default:
		panic(fmt.Sprintf("unsupported type: %v", ttype))
	}
}

// Not even sure we need this
func (c *TypeConverter) FetchNames() map[string]types.IType {
	allTypes := make(map[string]types.IType)
	for _, t := range c.seenTypes {
		names := c.GetTypeNames(t)
		for _, name := range names {
			_, exists := allTypes[name]
			if exists {
				continue
			}
			allTypes[name] = t
		}
	}
	return allTypes
}

func (tc *TypeConverter) convertPortableTypeV14(ttype substrateTypes.PortableTypeV14, allMetadataTypes []substrateTypes.PortableTypeV14) types.IType {
	lookupId := ttype.ID.Int64()
	if tc.seenTypes[lookupId] != nil {
		return tc.seenTypes[lookupId]
	}

	if ttype.Type.Def.IsComposite {
		c := ttype.Type.Def.Composite
		return tc.convertCompositeType(c, allMetadataTypes)
	}

	if ttype.Type.Def.IsVariant {
		v := ttype.Type.Def.Variant
		return tc.convertVariantType(v, allMetadataTypes)
	}

	if ttype.Type.Def.IsSequence {
		s := ttype.Type.Def.Sequence
		return tc.convertSequenceType(s, allMetadataTypes)
	}

	if ttype.Type.Def.IsArray {
		a := ttype.Type.Def.Array
		return tc.convertArrayType(a, allMetadataTypes)
	}

	if ttype.Type.Def.IsTuple {
		t := ttype.Type.Def.Tuple
		return tc.convertTupleType(t, allMetadataTypes)
	}

	if ttype.Type.Def.IsPrimitive {
		p := ttype.Type.Def.Primitive
		b := substrateTypes.Si0TypeDefPrimitive(p.Si0TypeDefPrimitive)
		return tc.convertPrimitiveType(b)
	}

	if ttype.Type.Def.IsCompact {
		c := ttype.Type.Def.Compact
		return tc.convertCompactType(c, allMetadataTypes)
	}

	if ttype.Type.Def.IsBitSequence {
		b := ttype.Type.Def.BitSequence
		return tc.convertBitSequenceType(b, allMetadataTypes)
	}

	if ttype.Type.Def.IsHistoricMetaCompat {
		h := ttype.Type.Def.HistoricMetaCompat
		return tc.convertHistoricMetaCompat(h)
	}

	panic(fmt.Sprintf("unsupported type: %v", ttype.Type.Def))
}

func (tc *TypeConverter) convertCompositeType(c substrateTypes.Si1TypeDefComposite, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Composite {
	comp := &types.Composite{}
	fields := make([]*types.CompositeField, 0)

	for _, field := range c.Fields {
		lookupId := field.Type.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)

		fields = append(fields, &types.CompositeField{
			Name: string(field.Name),
			Type: convertedType,
		})
	}

	comp.CompositeFields = fields
	return comp
}

type Optional struct {
	Type types.IType
}

func (o *Optional) ToProtoMessage(options ...string) string {
	return "optional " + o.Type.ToProtoMessage()
}

type Result struct {
	Type types.IType // error or types.Type
}

func (r *Result) ToProtoMessage(options ...string) string {
	return r.Type.ToProtoMessage()
}

type Error struct {
	Type types.IType
}

func (e *Error) ToProtoMessage(options ...string) string {
	return "Error"
}

type BaseType struct {
	Type types.IType
}

func (b *BaseType) ToProtoMessage(options ...string) string {
	return b.Type.ToProtoMessage()
}

func (b *BaseType) SetChild(c types.IType) {
	b.Type = c
}

// NOTE: inside a variant we can find the definition of a Rust Result type, which would need to be handled a bit funkier...
func (tc *TypeConverter) convertVariantType(v substrateTypes.Si1TypeDefVariant, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Variants {
	variant := &types.Variants{}
	variants := make([]*types.Variant, 0)

	for _, v := range v.Variants {
		fields := make([]*types.VariantField, 0)
		for _, field := range v.Fields {
			lookupId := field.Type.Int64()

			portableType := allMetadataTypes[lookupId]
			name := string(field.Name)

			// this is the only part that I am not sure...
			if name == "" {
				continue
			}

			tc.UpdateName(name)
			convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)

			fields = append(fields, &types.VariantField{
				Name: name,
				Type: convertedType,
			})
		}

		variants = append(variants, &types.Variant{
			Name:          string(v.Name),
			VariantFields: fields,
			Index:         uint64(v.Index) + 1,
		})
	}

	variant.Variants = variants
	return variant
}

func (tc *TypeConverter) convertSequenceType(s substrateTypes.Si1TypeDefSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Sequence {
	seq := &types.Sequence{}
	lookupId := s.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)
	seq.Item = convertedType
	return seq
}

func (tc *TypeConverter) convertArrayType(a substrateTypes.Si1TypeDefArray, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Array {
	arr := &types.Array{}
	lookupId := a.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)
	arr.Type = convertedType
	return arr
}

func (tc *TypeConverter) convertTupleType(t substrateTypes.Si1TypeDefTuple, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Tuple {
	tup := &types.Tuple{
		// Name: tc.name,
	}
	items := make([]types.IType, 0)

	for _, item := range t {
		lookupId := item.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)
		items = append(items, convertedType)
	}

	tup.Items = items
	return tup
}

func (tc *TypeConverter) convertPrimitiveType(b substrateTypes.Si0TypeDefPrimitive) *types.Primitive {
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

func (tc *TypeConverter) convertCompactType(c substrateTypes.Si1TypeDefCompact, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Compact {
	comp := &types.Compact{}
	lookupId := c.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := tc.convertPortableTypeV14(portableType, allMetadataTypes)
	comp.Type = convertedType
	return comp
}

func (tc *TypeConverter) convertBitSequenceType(b substrateTypes.Si1TypeDefBitSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.BitSequence {
	bitSeq := &types.BitSequence{}
	bitOrderPortableType := allMetadataTypes[b.BitOrderType.Int64()]
	bitStorePortableType := allMetadataTypes[b.BitStoreType.Int64()]
	bitSeq.BitOrderType = tc.convertPortableTypeV14(bitOrderPortableType, allMetadataTypes)
	bitSeq.BitStoreType = tc.convertPortableTypeV14(bitStorePortableType, allMetadataTypes)
	return bitSeq
}

func (tc *TypeConverter) convertHistoricMetaCompat(h substrateTypes.Type) types.HistoricMetaCompat {
	return types.HistoricMetaCompat(string(h))
}
