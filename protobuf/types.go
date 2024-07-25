package protobuf

import (
	"fmt"
	"strings"

	"github.com/streamingfast/firehose-gear/utils"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types2 "github.com/streamingfast/firehose-gear/types"
)

type Lookupable interface {
	GetLookupId() int64
}
type Field interface {
	SetOptional()
	IsPrimitive() bool
	IsRepeated() bool
	IsOneOf() bool
	IsOptional() bool
	FullTypeName() string
	GetType() string
	GetLookupId() int64
	ToGoTypeName(meta *types.Metadata) string
	ToProto(idx int) (string, int)
	ToFuncName(meta *types.Metadata) string
	ToFieldName() string
	ReturnType(meta *types.Metadata) string
	OutputType(meta *types.Metadata) string
}

type Proto struct {
	Syntax   string
	Version  string
	Imports  string
	Messages []*Message
}

type Message struct {
	Pallet   string
	Name     string
	Fields   []Field
	LookupID int64
}

func (m *Message) GetLookupId() int64 {
	return m.LookupID
}

func (m *Message) FullTypeName() string {
	suffix := utils.ToPascalCase(m.Name, utils.CapitalizeCharAfterNum)
	if m.Pallet == "" {
		return suffix
	}
	return utils.ToPascalCase(m.Pallet) + "_" + suffix
}
func (m *Message) ToFuncName(meta *types.Metadata) string {
	name := "To_" + m.FullTypeName()
	return name
}

func (m *Message) ReturnType(meta *types.Metadata) string {
	return "*pbgear." + m.FullTypeName()
}
func (m *Message) OutputType(meta *types.Metadata) string {
	return "&pbgear." + m.FullTypeName()
}

func (m *Message) ProtoMessageName() string {
	suffix := m.Name
	if m.Pallet == "" {
		return suffix
	}
	return m.Pallet + "_" + suffix
}

func (m *Message) ToProto() string {
	var sb strings.Builder

	sb.WriteString("message " + m.FullTypeName() + " {\n")
	idx := 0
	str := ""
	for _, field := range m.Fields {
		idx++
		str, idx = field.ToProto(idx)
		sb.WriteString("\t" + str + "\n")
	}
	sb.WriteString("}\n")

	return sb.String()
}

type BasicField struct {
	Optional  bool
	Primitive bool
	Pallet    string
	Type      string
	Name      string
	LookupID  int64
}

func (f *BasicField) GetLookupId() int64 {
	return f.LookupID
}

func (f *BasicField) OneOfWrapperOutputName() string {
	return fmt.Sprintf("&pbgear.%s_%s", f.Pallet, f.FullTypeName())
}

func (f *BasicField) ToChildFuncName(meta *types.Metadata) string {
	name := "To_" + f.FullTypeName()
	return name
}

func (f *BasicField) IsOptional() bool {
	return f.Optional
}

func (f *BasicField) IsRepeated() bool {
	return false
}

func (f *BasicField) IsOneOf() bool {
	return false
}

func (f *BasicField) ToFieldName() string {
	return utils.ToPascalCase(f.Name)
}

func (f *BasicField) ToGoTypeName(meta *types.Metadata) string {
	if f.Primitive {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return primitive.ToGoType()
	}
	return "pbgear." + f.FullTypeName()
}

func (f *BasicField) GetType() string {
	return f.Type
}
func (f *BasicField) FullTypeName() string {
	if f.Primitive {
		return f.Type
	}

	suffix := utils.ToPascalCase(f.Type, utils.CapitalizeCharAfterNum)
	if f.Pallet == "" {
		return suffix
	}

	return utils.ToPascalCase(f.Pallet) + "_" + suffix
}

func (f *BasicField) ToFuncName(meta *types.Metadata) string {
	if f.Optional {
		return "To_optional_" + f.FullTypeName()
	}
	name := "To_" + f.FullTypeName()
	return name
}

func (f *BasicField) ReturnType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return primitive.ToGoType()
	}
	return "*pbgear." + f.FullTypeName()
}

func (f *BasicField) OutputType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return primitive.ToGoType()
	}
	return "&pbgear." + f.FullTypeName()
}

func (f *BasicField) SetOptional() {
	f.Optional = true
}

func (f *BasicField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("%s %s = %d;", f.FullTypeName(), utils.ToSnakeCase(f.Name), idx)
	if f.Optional {
		str = "optional " + str
	}
	return str, idx
}

func (f *BasicField) IsPrimitive() bool {
	return f.Primitive
}

type RepeatedField struct {
	Pallet    string
	Type      string
	Name      string
	LookupID  int64
	Primitive bool
}

func (f *RepeatedField) GetLookupId() int64 {
	return f.LookupID
}

func (f *RepeatedField) IsOptional() bool {
	return false
}

func (f *RepeatedField) IsRepeated() bool {
	return true
}

func (f *RepeatedField) IsOneOf() bool {
	return false
}

func (f *RepeatedField) ToFieldName() string {
	return utils.ToPascalCase(f.Name)
}

func (f *RepeatedField) ToGoTypeName(meta *types.Metadata) string {
	if f.Primitive {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return primitive.ToGoType()
	}

	return "pbgear." + f.FullTypeName()
}

func (f *RepeatedField) GetType() string {
	return f.Type
}

func (f *RepeatedField) FullTypeName() string {
	if f.Primitive {
		return f.Type
	}

	suffix := utils.ToPascalCase(f.Type, utils.CapitalizeCharAfterNum)
	if f.Pallet == "" {
		return suffix
	}
	return utils.ToPascalCase(f.Pallet) + "_" + suffix
}
func (f *RepeatedField) ToChildFuncName(meta *types.Metadata) string {
	name := "To_" + f.FullTypeName()
	//if name == "To_GprimitivesActorId" {
	//	panic("To_GprimitivesActorId")
	//}

	return name
}

func (f *RepeatedField) ToFuncName(meta *types.Metadata) string {
	return "To_repeated_" + f.FullTypeName()
}

func (f *RepeatedField) ReturnType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return "[]" + primitive.ToGoType()
	}
	return "[]*pbgear." + f.FullTypeName()

}
func (f *RepeatedField) OutputType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return "[]" + primitive.ToGoType()
	}
	return "[]*pbgear." + f.FullTypeName()

}

func (f *RepeatedField) SetOptional() {
	panic("Repeated fields can not be set with an optional")
}

func (f *RepeatedField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("repeated %s %s = %d;", f.FullTypeName(), utils.ToSnakeCase(f.Name), idx)
	return str, idx
}

func (f *RepeatedField) IsPrimitive() bool {
	return f.Primitive
}

type OneOfField struct {
	Pallet    string
	Name      string
	Types     []*BasicField
	LookupID  int64
	Primitive bool
}

func (f *OneOfField) GetLookupId() int64 {
	return f.LookupID
}

func (f *OneOfField) Toto() string {
	return "foobar"
}

func (f *OneOfField) IsOptional() bool {
	return false
}

func (f *OneOfField) ToFieldName() string {
	return utils.ToPascalCase(f.Name)
}

func (f *OneOfField) IsOneOf() bool {
	return true
}

func (f *OneOfField) IsRepeated() bool {
	return false
}

func (f *OneOfField) ToGoTypeName(meta *types.Metadata) string {
	return "pbgear." + f.FullTypeName()
}

func (f *OneOfField) GetType() string {
	panic("OneOfField has not type")
}
func (f *OneOfField) FullTypeName() string {
	suffix := utils.ToPascalCase(f.Name, utils.CapitalizeCharAfterNum)
	if f.Pallet == "" {
		return suffix
	}
	return utils.ToPascalCase(f.Pallet) + "_" + suffix
}

func (f *OneOfField) ToFuncName(meta *types.Metadata) string {
	return "To_oneof_" + f.FullTypeName()
}

func (f *OneOfField) ReturnType(meta *types.Metadata) string {
	return "*pbgear." + f.FullTypeName()
}

func (f *OneOfField) OutputType(meta *types.Metadata) string {
	return "&pbgear." + f.FullTypeName()
}

func (f *OneOfField) ToProto(idx int) (string, int) {
	var sb strings.Builder

	sb.WriteString("oneof " + f.Name + " {\n")
	for _, field := range f.Types {
		sb.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", field.FullTypeName(), utils.ToSnakeCase(field.Name), idx))
		idx++
	}
	sb.WriteString("\t}")
	idx--
	return sb.String(), idx
}

func (f *OneOfField) SetOptional() {
	panic("One field can not be set with an optional")
}

func (f *OneOfField) IsPrimitive() bool {
	return f.Primitive
}
