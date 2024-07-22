package protobuf

import (
	"fmt"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types2 "github.com/streamingfast/firehose-gear/types"
)

type Field interface {
	GetType() string
	SetOptional()
	IsPrimitive() bool
	ToProto(idx int) (string, int)
	ToFuncName(meta *types.Metadata) string
	ReturnType(meta *types.Metadata) string
}

type Proto struct {
	Syntax   string
	Version  string
	Imports  string
	Messages []*Message
	Enums    []*Enums
}

type Enums struct {
}

type Message struct {
	Name   string
	Fields []Field
}

func (m *Message) ToFuncName(meta *types.Metadata) string {
	return "to_" + m.Name
}

func (m *Message) ReturnType(meta *types.Metadata) string {
	return "*pbgear." + m.Name
}

func (m *Message) ToProto() string {
	var sb strings.Builder

	sb.WriteString("message " + m.Name + " {\n")
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
	Type      string
	Name      string
	LookupID  int64
	Primitive bool
}

func (f *BasicField) ToFuncName(meta *types.Metadata) string {
	if f.Optional {
		return "to_optional_" + f.Name
	}
	return "to_" + f.Name
}

func (f *BasicField) ReturnType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return primitive.ToGoType()
	}
	return "*pbgear." + f.Name
}

func (f *BasicField) GetType() string {
	return f.Type
}

func (f *BasicField) SetOptional() {
	f.Optional = true
}

func (f *BasicField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("%s %s = %d;", f.Type, f.Name, idx)
	if f.Optional {
		str = "optional " + str
	}
	return str, idx
}

func (f *BasicField) IsPrimitive() bool {
	return f.Primitive
}

type RepeatedField struct {
	Type      string
	Name      string
	LookupID  int64
	Primitive bool
}

func (f *RepeatedField) ToFuncName() string {
	return "to_repeated" + f.Name
}

func (f *RepeatedField) ReturnType(meta *types.Metadata) string {
	if f.IsPrimitive() {
		ttype := meta.AsMetadataV14.EfficientLookup[f.LookupID]
		primitive := types2.ConvertPrimitiveType(ttype.Def.Primitive.Si0TypeDefPrimitive)
		return "[]" + primitive.ToGoType()
	}
	return "[]*pbgear." + f.Name

}

func (f *RepeatedField) SetOptional() {
	panic("Repeated fields can not be set with an optional")
}

func (f *RepeatedField) GetType() string {
	return f.Type
}

func (f *RepeatedField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("repeated %s %s = %d;\n", f.Type, f.Name, idx)
	return str, idx
}

func (f *RepeatedField) IsPrimitive() bool {
	return f.Primitive
}

type OneOfField struct {
	Name      string
	Types     []*BasicField
	LookupID  int64
	Primitive bool
}

func (f *OneOfField) ToFuncName() string {
	return "to_oneof_" + f.Name
}

func (f *OneOfField) ReturnType(meta *types.Metadata) string {
	return "*pbgear." + f.Name
}

func (f *OneOfField) GetType() string {
	panic("OneOfField does not have a type")
}

func (f *OneOfField) ToProto(idx int) (string, int) {
	var sb strings.Builder

	sb.WriteString("oneof " + f.Name + " {\n")
	for _, field := range f.Types {
		sb.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", field.Type, field.Name, idx))
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
