package protobuf

import (
	"fmt"
	"strings"
)

type Field interface {
	GetType() string
	SetOptional()
	IsPrimitive() bool
	ToProto(idx int) (string, int)
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

type RepeatedField struct {
	Type      string
	Name      string
	Primitive bool
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

type BasicField struct {
	Optional  bool
	Type      string
	Name      string
	Primitive bool
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

type OneOfField struct {
	Name      string
	Types     []*BasicField
	Primitive bool
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
