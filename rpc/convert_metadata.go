package rpc

import (
	"fmt"
	"math"
	"os"
	"strings"

	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
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

	tc       *TypeConverter
	metaData *substrateTypes.Metadata
}

func NewMetadataConverter(gearClients *firecoreRPC.Clients[*Client], logger *zap.Logger, tracer logging.Tracer) *MetadataConverter {
	return &MetadataConverter{
		gearClients: gearClients,
		logger:      logger,
		tracer:      tracer,
	}
}

func (mc *MetadataConverter) FetchMessages() []*protobuf.Message {
	outputs := make([]*protobuf.Message, 0)
	for _, seenMsg := range mc.tc.messages {
		if seenMsg != nil {
			outputs = append(outputs, seenMsg)
		}
	}
	return outputs
}

func (mc *MetadataConverter) FetchMetadata() *substrateTypes.Metadata {
	return mc.metaData
}

func (mc *MetadataConverter) Convert(blockHash string) (map[string]types.IType, error) {
	metadata, err := mc.fetchStateMetadata(blockHash)
	if err != nil {
		return nil, fmt.Errorf("fetching state metadata: %w", err)
	}

	mc.metaData = metadata
	if metadata.Version != 14 {
		return nil, nil
	}

	switch metadata.Version {
	case 14:
		mc.tc = &TypeConverter{
			messages:         make(map[string]*protobuf.Message),
			allMetadataTypes: []substrateTypes.PortableTypeV14{},
		}
		return mc.tc.convertTypesFromv14(metadata.AsMetadataV14)
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
	messages         map[string]*protobuf.Message
	allMetadataTypes []substrateTypes.PortableTypeV14
}

func (c *TypeConverter) convertTypesFromv14(metadata substrateTypes.MetadataV14) (map[string]types.IType, error) {
	allMetadataTypes := metadata.Lookup.Types
	c.allMetadataTypes = allMetadataTypes

	for _, pallet := range metadata.Pallets {
		//if pallet.Name != "Babe" {
		//	continue
		//}
		if pallet.HasCalls {
			callIdx := pallet.Calls.Type.Int64()
			palletName := string(pallet.Name)
			c.ProcessPalletCalls(callIdx, palletName)
		}

		// if pallet.HasEvents {
		// 	idx := pallet.Events.Type.Int64()
		// 	variants := c.allMetadataTypes[idx]

		// 	events := variants.Type.Def.Variant
		// 	totalNumberOfCalls += len(events.Variants)
		// 	for _, variant := range events.Variants {
		// 		// if variant.Name != "as_multi_threshold_1" {
		// 		// 	continue
		// 		// }
		// 		// fmt.Println("Processing call", variant.Name)

		// 		palletName := string(pallet.Name)
		// 		callName := string(variant.Name)
		// 		messageName := stringy.New(palletName).PascalCase().Get() + "_" + stringy.New(callName).PascalCase().Get() + "_Call"
		// 		message := &protobuf.Message{
		// 			Name: messageName,
		// 		}

		// 		c.ProcessCallFields(variant, message, palletName, callName)
		// 		c.seenMessages[message.Name] = message
		// 	}
		// }
	}

	outputs := make([]*protobuf.Message, 0)
	for _, seenMsg := range c.messages {
		if seenMsg != nil {
			outputs = append(outputs, seenMsg)
		}
	}

	var sb strings.Builder
	sb.WriteString("syntax = \"proto3\";\n")
	sb.WriteString("package sf.gear.metadata.type.v1;\n")
	sb.WriteString("option go_package = \"github.com/streamingfast/firehose-gear/pb/sf/gear/metadata/type/v1;pbgear\";\n\n")
	for _, out := range outputs {
		sb.WriteString(out.ToProto())
	}
	err := os.WriteFile("../proto/sf/gear/metadata/type/v1/output.proto", []byte(sb.String()), 0644)
	if err != nil {
		return nil, fmt.Errorf("writing output.proto: %w", err)
	}

	return nil, nil
}

func (c *TypeConverter) ProcessPalletCalls(callIdx int64, palletName string) {
	variants := c.allMetadataTypes[callIdx]

	calls := variants.Type.Def.Variant
	for _, variant := range calls.Variants {
		//if variant.Name != "report_equivocation_unsigned" { //lowercase
		//	continue
		//}
		callName := string(variant.Name)
		message := &protobuf.Message{
			Pallet: palletName,
			Name:   callName + "_Call",
		}

		c.ProcessCallFields(variant, message, palletName, callName)
		c.messages[message.FullTypeName()] = message
	}
}

func (c *TypeConverter) ProcessCallFields(variant substrateTypes.Si1Variant, message *protobuf.Message, palletName string, callName string) {
	for _, f := range variant.Fields {
		fieldName := string(f.Name)
		field := c.ProcessField(f, palletName, callName, fieldName)
		if field != nil {
			message.Fields = append(message.Fields, field)
		}
	}
}

func (c *TypeConverter) ProcessField(f substrateTypes.Si1Field, palletName string, callName string, fieldName string) protobuf.Field {
	idx := f.Type.Int64()
	ttype := c.allMetadataTypes[idx]

	return c.FieldForType(ttype, palletName, callName, fieldName)
}

func (c *TypeConverter) FieldForPrimitive(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	return &protobuf.BasicField{
		Pallet:    palletName,
		Name:      fieldName,
		Type:      types.ConvertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).GetProtoFieldName(),
		Primitive: true,
		LookupID:  ttype.ID.Int64(),
	}
}

func (c *TypeConverter) FieldForSequence(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.RepeatedField {
	lookupId := ttype.Type.Def.Sequence.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	return &protobuf.RepeatedField{
		Pallet:    palletName,
		Type:      typeName,
		Name:      fieldName,
		Primitive: lookupType.Type.Def.IsPrimitive,
		LookupID:  lookupId,
	}
}

func (c *TypeConverter) FieldForArray(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.RepeatedField {
	lookupId := ttype.Type.Def.Array.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	return &protobuf.RepeatedField{
		Pallet:    palletName,
		Name:      fieldName,
		Type:      typeName,
		Primitive: lookupType.Type.Def.IsPrimitive,
		LookupID:  lookupId,
	}
}

func (c *TypeConverter) FieldForCompact(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	name := c.ExtractTypeName(ttype, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Pallet:   palletName,
		Name:     fieldName,
		Type:     name,
		LookupID: ttype.ID.Int64(),
	}
}

func (c *TypeConverter) FieldForComposite(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	name := c.ExtractTypeName(ttype, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Pallet:   palletName,
		Name:     fieldName,
		Type:     name,
		LookupID: ttype.ID.Int64(),
	}
}

func (c *TypeConverter) FieldForType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	if ttype.ID.Int64() == 65 {
		return c.FieldFor65(ttype, palletName, callName, fieldName)
	}

	var field protobuf.Field
	var optional bool
	if len(ttype.Type.Path) != 0 {
		isOptional := ttype.Type.Path[len(ttype.Type.Path)-1]
		optional = isOptional == "Option"
	}

	if optional {
		lookupId := ttype.Type.Def.Variant.Variants[1].Fields[0].Type.Int64()
		ttype = c.allMetadataTypes[lookupId]
	}

	switch {
	case ttype.Type.Def.IsPrimitive:
		field = c.FieldForPrimitive(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}
		return field

	case ttype.Type.Def.IsSequence:
		field = c.FieldForSequence(ttype, palletName, callName, fieldName)
		lookupId := ttype.Type.Def.Sequence.Type.Int64()
		t := c.allMetadataTypes[lookupId]
		if t.Type.Def.IsTuple {
			ttype = t
		}

	case ttype.Type.Def.IsArray:
		field = c.FieldForArray(ttype, palletName, callName, fieldName)
		lookupId := ttype.Type.Def.Array.Type.Int64()
		t := c.allMetadataTypes[lookupId]
		if t.Type.Def.IsTuple {
			ttype = t
		}

	case ttype.Type.Def.IsTuple:
		field = c.FieldForTuple(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsVariant:
		field = c.FieldForVariant(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsCompact:
		field = c.FieldForCompact(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsComposite:
		field = c.FieldForComposite(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}
	}

	if !ttype.Type.Def.IsVariant {
		if _, ok := c.messages[field.FullTypeName()]; !ok {
			if field.IsPrimitive() {
				c.messages[field.FullTypeName()] = nil
			} else {
				msg := c.MessageForType(field.GetType(), ttype, palletName, callName, fieldName)
				c.messages[field.FullTypeName()] = msg
			}
		}
	}

	return field
}

func (c *TypeConverter) ExtractTypeName(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	typeName := ""
	if ttype.Type.Def.IsPrimitive {
		return types.ConvertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).ToProtoType()
	}

	if ttype.Type.Def.IsTuple {
		typeName = c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsComposite {
		typeName = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsVariant {
		typeName = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsSequence || ttype.Type.Def.IsArray {
		typeName = fmt.Sprintf("%s_%s_list", palletName, fieldName)
	}

	if ttype.Type.Def.IsCompact {
		lookupId := ttype.Type.Def.Compact.Type.Int64()
		lookupType := c.allMetadataTypes[lookupId]
		typeName = c.ExtractTypeName(lookupType, palletName, callName, fieldName)
		typeName = fmt.Sprintf("Compact_%s", typeName)
	}

	return typeName
}

func (c *TypeConverter) MessageForType(typeName string, ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.Message {
	msg := &protobuf.Message{
		Pallet: palletName,
		Name:   typeName,
	}

	if ttype.Type.Def.IsTuple {
		if typeName != "Tuple_Null" {

			for i, idx := range ttype.Type.Def.Tuple {
				lookupId := idx.Int64()
				t := c.allMetadataTypes[lookupId]
				f := c.FieldForType(t, palletName, callName, fmt.Sprintf("value_%d", i))
				msg.Fields = append(msg.Fields, f)
			}
		}
	}

	if ttype.Type.Def.IsComposite {
		if len(ttype.Type.Def.Composite.Fields) == 1 {
			field := ttype.Type.Def.Composite.Fields[0]
			lookup := field.Type.Int64()
			lookupType := c.allMetadataTypes[lookup]

			fName := string(field.Name)
			if !field.HasName {
				fName = fieldName
			}

			f := c.FieldForType(lookupType, palletName, callName, fName)
			msg.Fields = append(msg.Fields, f)
		} else {
			for _, f := range ttype.Type.Def.Composite.Fields {
				lookup := f.Type.Int64()
				lookupType := c.allMetadataTypes[lookup]
				fieldName := string(f.Name)
				field := c.FieldForType(lookupType, palletName, callName, fieldName)
				msg.Fields = append(msg.Fields, field)
			}
		}
	}

	if ttype.Type.Def.IsVariant {
		field := &protobuf.OneOfField{
			Pallet:   palletName,
			Name:     typeName,
			LookupID: ttype.ID.Int64(),
		}
		for _, v := range ttype.Type.Def.Variant.Variants {
			field.Types = append(field.Types, &protobuf.BasicField{
				Pallet:   palletName,
				Name:     fmt.Sprintf("%s_%s", palletName, string(v.Name)),
				Type:     string(v.Name),
				LookupID: math.MaxInt64,
			})
		}
		msg.Fields = append(msg.Fields, field)
	}

	if ttype.Type.Def.IsSequence {
		lookupId := ttype.Type.Def.Sequence.Type.Int64()
		childType := c.allMetadataTypes[lookupId]

		f := c.FieldForType(childType, palletName, callName, fieldName)
		msg.Fields = append(msg.Fields, f)
	}

	if ttype.Type.Def.IsArray {
		lookupId := ttype.Type.Def.Array.Type.Int64()
		childType := c.allMetadataTypes[lookupId]

		f := c.FieldForType(childType, palletName, callName, fieldName)
		msg.Fields = append(msg.Fields, f)
	}

	if ttype.Type.Def.IsCompact {
		lookupId := ttype.Type.Def.Compact.Type.Int64()
		childType := c.allMetadataTypes[lookupId]

		field := c.FieldForType(childType, palletName, callName, "value")

		msg.Fields = append(msg.Fields, field)
	}

	return msg
}

func (c *TypeConverter) FieldFor65(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	of := &protobuf.OneOfField{
		Pallet:   palletName,
		Name:     fieldName,
		LookupID: ttype.ID.Int64(),
	}

	pallets := ttype.Type.Def.Variant.Variants
	for _, v := range pallets { // 1. System
		palletName := v.Name
		calls := v.Fields
		for _, call := range calls { // 1. 66
			lookupId := call.Type.Int64()
			palletCalls := c.allMetadataTypes[lookupId] // 1. System
			palletCallNames := palletCalls.Type.Def.Variant.Variants

			for _, palletCallName := range palletCallNames { // remark
				of.Types = append(of.Types, &protobuf.BasicField{
					Pallet:   string(palletName),
					Name:     fmt.Sprintf("%s_%s", palletName, string(palletCallName.Name)),
					LookupID: math.MaxInt64,
					Type:     string(palletCallName.Name) + "_Call",
				})
			}
		}
	}

	return of
}

func (c *TypeConverter) MessageForVariantTypes(name string, variant substrateTypes.Si1Variant, palletName string, callName string, fieldName string) {
	msg := &protobuf.Message{
		Pallet: palletName,
		Name:   name,
	}

	for i, f := range variant.Fields {
		idx := f.Type.Int64()
		fieldType := c.allMetadataTypes[idx]
		fn := string(f.Name)
		if !f.HasName {
			fn = fmt.Sprintf("value_%d", i)
		}
		field := c.FieldForType(fieldType, palletName, callName, fn)
		msg.Fields = append(msg.Fields, field)
	}

	c.messages[msg.FullTypeName()] = msg
}

func (c *TypeConverter) ExtractTypeNameFromTuple(tuple substrateTypes.Si1TypeDefTuple, palletName string, callName string, fieldName string) string {
	var name strings.Builder
	name.WriteString("Tuple_")
	if len(tuple) == 0 {
		name.WriteString("Null")
	}

	for _, item := range tuple {
		lookupId := item.Int64()
		ttype := c.allMetadataTypes[lookupId]
		name.WriteString(c.ExtractTypeName(ttype, palletName, callName, fieldName))
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

func (c *TypeConverter) FieldForTuple(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	if fieldName == "" {
		panic("field name is required")
	}

	field := &protobuf.BasicField{
		Pallet:   palletName,
		Name:     fieldName,
		Type:     c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple, palletName, callName, fieldName),
		LookupID: ttype.ID.Int64(),
	}

	return field
}

func (c *TypeConverter) FieldForVariant(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {

	f := &protobuf.BasicField{
		Pallet:   palletName,
		Name:     fieldName,
		Type:     fieldName,
		LookupID: ttype.ID.Int64(),
	}

	oneOf := &protobuf.OneOfField{
		Name:   "value",
		Pallet: palletName,
	}
	msg := &protobuf.Message{
		Pallet: palletName,
		Name:   fieldName,
		Fields: []protobuf.Field{oneOf},
	}

	if _, found := c.messages[msg.FullTypeName()]; found {
		return f
	}

	for _, v := range ttype.Type.Def.Variant.Variants {
		typeName := string(v.Name)
		oneOf.Types = append(oneOf.Types, &protobuf.BasicField{
			Pallet:   palletName,
			Name:     string(v.Name),
			Type:     string(v.Name),
			LookupID: math.MaxInt64,
		})

		if _, ok := c.messages[typeName]; !ok {
			c.MessageForVariantTypes(typeName, v, palletName, callName, string(v.Name))
		}
	}

	c.messages[msg.FullTypeName()] = msg
	return f
}

func (c *TypeConverter) ProcessOptionalType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	variant := ttype.Type.Def.Variant
	some := variant.Variants[1]
	someField := some.Fields[0]
	tttype := c.allMetadataTypes[someField.Type.Int64()]

	return c.ExtractTypeName(tttype, palletName, callName, fieldName)
}
