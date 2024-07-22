package rpc

import (
	"fmt"
	"os"
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
	types            map[string]types.IType
	seenTypes        map[int64]types.IType
	seenMessages     map[string]*protobuf.Message
	allMetadataTypes []substrateTypes.PortableTypeV14
}

func (c *TypeConverter) UpdateName(name string) {
	// tc.name = name
}

func (c *TypeConverter) convertTypesFromv14(metadata substrateTypes.MetadataV14) (map[string]types.IType, error) {
	allMetadataTypes := metadata.Lookup.Types
	c.allMetadataTypes = allMetadataTypes
	totalNumberOfCalls := 0

	for _, pallet := range metadata.Pallets {
		//if pallet.Name != "GearVoucher" {
		//	continue
		//}
		if pallet.HasCalls {
			idx := pallet.Calls.Type.Int64()
			variants := c.allMetadataTypes[idx]

			calls := variants.Type.Def.Variant
			totalNumberOfCalls += len(calls.Variants)
			for _, variant := range calls.Variants {
				//if variant.Name != "issue" {
				//	continue
				//}
				fmt.Println("Processing call", variant.Name)

				palletName := string(pallet.Name)
				callName := string(variant.Name)
				messageName := stringy.New(palletName).PascalCase().Get() + "_" + stringy.New(callName).PascalCase().Get() + "_Call"
				message := &protobuf.Message{
					Name: messageName,
				}

				c.ProcessCallFields(variant, message, palletName, callName)
				c.seenMessages[message.Name] = message
			}
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

	for _, seenMsg := range c.seenMessages {
		if seenMsg != nil {
			outputs = append(outputs, seenMsg)
		}
	}

	var sb strings.Builder
	sb.WriteString("syntax = \"proto3\";\n\n")
	for _, out := range outputs {
		sb.WriteString(out.ToProto())
	}
	err := os.WriteFile("../proto/output.proto", []byte(sb.String()), 0644)
	if err != nil {
		return nil, fmt.Errorf("writing output.proto: %w", err)
	}

	return nil, nil
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

func (c *TypeConverter) FieldForPrimitive(ttype substrateTypes.PortableTypeV14, fieldName string) *protobuf.BasicField {
	return &protobuf.BasicField{
		Name:      fieldName,
		Type:      c.convertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).GetProtoFieldName(),
		Primitive: true,
	}
}

func (c *TypeConverter) FieldForSequence(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.RepeatedField {
	lookupId := ttype.Type.Def.Sequence.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	return &protobuf.RepeatedField{
		Type:      typeName,
		Name:      fieldName,
		Primitive: lookupType.Type.Def.IsPrimitive,
	}
}

func (c *TypeConverter) FieldForArray(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.RepeatedField {
	lookupId := ttype.Type.Def.Array.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	return &protobuf.RepeatedField{
		Name:      fieldName,
		Type:      typeName,
		Primitive: lookupType.Type.Def.IsPrimitive,
	}
}

func (c *TypeConverter) FieldForCompact(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	name := c.ExtractTypeName(ttype, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Name: fieldName,
		Type: name,
	}
}

func (c *TypeConverter) FieldForComposite(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	name := c.ExtractTypeName(ttype, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Name: fieldName,
		Type: name,
	}
}

func (c *TypeConverter) FieldForType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	if ttype.ID.Int64() == 65 {
		return c.FieldFor65(ttype, palletName, callName, fieldName)
	}

	var field protobuf.Field
	if len(ttype.Type.Path) != 0 {
		isOptional := ttype.Type.Path[len(ttype.Type.Path)-1]
		if isOptional == "Option" {
			optionType := c.ProcessOptionalType(ttype, palletName, callName, fieldName)
			field = &protobuf.BasicField{
				Optional:  true,
				Name:      fieldName,
				Type:      optionType,
				Primitive: ttype.Type.Def.IsPrimitive,
			}
		}
	}

	if ttype.Type.Def.IsPrimitive {
		return c.FieldForPrimitive(ttype, fieldName)
	}

	if ttype.Type.Def.IsSequence {
		field = c.FieldForSequence(ttype, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsArray {
		field = c.FieldForArray(ttype, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsTuple {
		field = c.FieldForTuple(ttype, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsVariant {
		field = c.FieldForVariant(ttype, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsCompact {
		field = c.FieldForCompact(ttype, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsComposite {
		field = c.FieldForComposite(ttype, palletName, callName, fieldName)
	}

	if !ttype.Type.Def.IsVariant {
		if _, ok := c.seenMessages[field.GetType()]; !ok {
			if field.IsPrimitive() {
				c.seenMessages[field.GetType()] = nil
			} else {
				msg := c.MessageForType(field.GetType(), ttype, palletName, callName, fieldName)
				c.seenMessages[field.GetType()] = msg
			}
		}
	}

	return field
}

func (c *TypeConverter) ExtractTypeName(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	typeName := ""
	if ttype.Type.Def.IsPrimitive {
		typeName = c.convertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).ToProtoMessage()
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
		typeName = "Compact"

		lookupId := ttype.Type.Def.Compact.Type.Int64()
		lookupType := c.allMetadataTypes[lookupId]
		typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)
		typeName = fmt.Sprintf("%s_%s", typeName, typeName)
	}

	return typeName
}

func (c *TypeConverter) MessageForType(typeName string, ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.Message {
	msg := &protobuf.Message{
		Name: typeName, // + "_GEN",
	}

	if ttype.Type.Def.IsTuple {
		msg.Fields = append(msg.Fields, c.FieldForTuple(ttype, palletName, callName, fieldName))
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
			Name: typeName,
		}
		for _, v := range ttype.Type.Def.Variant.Variants {
			typeName := fmt.Sprintf("%s_%s", palletName, string(v.Name))
			field.Types = append(field.Types, &protobuf.BasicField{
				Name: fmt.Sprintf("%s_%s", palletName, string(v.Name)),
				Type: typeName,
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
		typeName := c.ExtractTypeName(childType, palletName, callName, fieldName)

		field := &protobuf.BasicField{
			Name:      "value",
			Type:      typeName,
			Primitive: ttype.Type.Def.IsPrimitive,
		}

		msg.Fields = append(msg.Fields, field)
	}

	return msg
}

func (c *TypeConverter) FieldFor65(ttype substrateTypes.PortableTypeV14, _ string, callName string, fieldName string) protobuf.Field {
	of := &protobuf.OneOfField{
		Name: fieldName,
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
				n := string(palletName) + "_"
				n += stringy.New(string(palletCallName.Name)).PascalCase().Get()
				n += "_Call"
				of.Types = append(of.Types, &protobuf.BasicField{
					Name: fmt.Sprintf("%s_%s", palletName, string(palletCallName.Name)),
					Type: n,
				})
			}
		}
	}

	return of
}

func (c *TypeConverter) MessageForVariantTypes(name string, variant substrateTypes.Si1Variant, palletName string, callName string, fieldName string) {
	msg := &protobuf.Message{
		Name: name,
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

	c.seenMessages[msg.Name] = msg
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
		Name: fieldName,
		Type: c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple, palletName, callName, fieldName),
	}

	return field
}

func (c *TypeConverter) FieldForVariant(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	msgName := palletName + "_" + fieldName

	f := &protobuf.BasicField{
		Name: fieldName,
		Type: msgName,
	}

	if _, found := c.seenMessages[msgName]; found {
		return f
	}

	oneOf := &protobuf.OneOfField{
		Name: "value",
	}
	msg := &protobuf.Message{
		Name:   msgName,
		Fields: []protobuf.Field{oneOf},
	}

	for _, v := range ttype.Type.Def.Variant.Variants {

		typeName := fmt.Sprintf("%s_%s", palletName, string(v.Name))
		oneOf.Types = append(oneOf.Types, &protobuf.BasicField{
			Name: fmt.Sprintf("%s_%s", palletName, string(v.Name)),
			Type: typeName,
		})

		if _, ok := c.seenMessages[typeName]; !ok {
			c.MessageForVariantTypes(typeName, v, palletName, callName, string(v.Name))
		}
	}

	c.seenMessages[msgName] = msg

	return f
}

func (c *TypeConverter) ProcessOptionalType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	variant := ttype.Type.Def.Variant
	some := variant.Variants[1]
	someField := some.Fields[0]
	tttype := c.allMetadataTypes[someField.Type.Int64()]

	return c.ExtractTypeName(tttype, palletName, callName, fieldName)
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

func (c *TypeConverter) convertPortableTypeV14(ttype substrateTypes.PortableTypeV14, allMetadataTypes []substrateTypes.PortableTypeV14) types.IType {
	lookupId := ttype.ID.Int64()
	if c.seenTypes[lookupId] != nil {
		return c.seenTypes[lookupId]
	}

	if ttype.Type.Def.IsComposite {
		return c.convertCompositeType(ttype.Type.Def.Composite, allMetadataTypes)
	}

	if ttype.Type.Def.IsVariant {
		return c.convertVariantType(ttype.Type.Def.Variant, allMetadataTypes)
	}

	if ttype.Type.Def.IsSequence {
		return c.convertSequenceType(ttype.Type.Def.Sequence, allMetadataTypes)
	}

	if ttype.Type.Def.IsArray {
		return c.convertArrayType(ttype.Type.Def.Array, allMetadataTypes)
	}

	if ttype.Type.Def.IsTuple {
		return c.convertTupleType(ttype.Type.Def.Tuple, allMetadataTypes)
	}

	if ttype.Type.Def.IsPrimitive {
		b := ttype.Type.Def.Primitive.Si0TypeDefPrimitive
		return c.convertPrimitiveType(b)
	}

	if ttype.Type.Def.IsCompact {
		return c.convertCompactType(ttype.Type.Def.Compact, allMetadataTypes)
	}

	if ttype.Type.Def.IsBitSequence {
		return c.convertBitSequenceType(ttype.Type.Def.BitSequence, allMetadataTypes)
	}

	if ttype.Type.Def.IsHistoricMetaCompat {
		return c.convertHistoricMetaCompat(ttype.Type.Def.HistoricMetaCompat)
	}

	panic(fmt.Sprintf("unsupported type: %v", ttype.Type.Def))
}

func (c *TypeConverter) convertCompositeType(composite substrateTypes.Si1TypeDefComposite, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Composite {
	comp := &types.Composite{}
	fields := make([]*types.CompositeField, 0)

	for _, field := range composite.Fields {
		lookupId := field.Type.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)

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
func (c *TypeConverter) convertVariantType(v substrateTypes.Si1TypeDefVariant, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Variants {
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

			c.UpdateName(name)
			convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)

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

func (c *TypeConverter) convertSequenceType(s substrateTypes.Si1TypeDefSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Sequence {
	seq := &types.Sequence{}
	lookupId := s.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)
	seq.Item = convertedType
	return seq
}

func (c *TypeConverter) convertArrayType(a substrateTypes.Si1TypeDefArray, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Array {
	arr := &types.Array{}
	lookupId := a.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)
	arr.Type = convertedType
	return arr
}

func (c *TypeConverter) convertTupleType(t substrateTypes.Si1TypeDefTuple, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Tuple {
	tup := &types.Tuple{
		// Name: tc.name,
	}
	items := make([]types.IType, 0)

	for _, item := range t {
		lookupId := item.Int64()
		portableType := allMetadataTypes[lookupId]
		convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)
		items = append(items, convertedType)
	}

	tup.Items = items
	return tup
}

func (c *TypeConverter) convertPrimitiveType(b substrateTypes.Si0TypeDefPrimitive) *types.Primitive {
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

func (c *TypeConverter) convertCompactType(compact substrateTypes.Si1TypeDefCompact, allMetadataTypes []substrateTypes.PortableTypeV14) *types.Compact {
	comp := &types.Compact{}
	lookupId := compact.Type.Int64()
	portableType := allMetadataTypes[lookupId]
	convertedType := c.convertPortableTypeV14(portableType, allMetadataTypes)
	comp.Type = convertedType
	return comp
}

func (c *TypeConverter) convertBitSequenceType(b substrateTypes.Si1TypeDefBitSequence, allMetadataTypes []substrateTypes.PortableTypeV14) *types.BitSequence {
	bitSeq := &types.BitSequence{}
	bitOrderPortableType := allMetadataTypes[b.BitOrderType.Int64()]
	bitStorePortableType := allMetadataTypes[b.BitStoreType.Int64()]
	bitSeq.BitOrderType = c.convertPortableTypeV14(bitOrderPortableType, allMetadataTypes)
	bitSeq.BitStoreType = c.convertPortableTypeV14(bitStorePortableType, allMetadataTypes)
	return bitSeq
}

func (c *TypeConverter) convertHistoricMetaCompat(h substrateTypes.Type) types.HistoricMetaCompat {
	return types.HistoricMetaCompat(string(h))
}
