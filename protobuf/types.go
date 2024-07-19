package protobuf

import (
	"fmt"
	"strings"
)

type Field interface {
	GetType() string
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

func (r *RepeatedField) GetType() string {
	return r.Type
}

func (r *RepeatedField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("repeated %s %s = %d;\n", r.Type, r.Name, idx)
	return str, idx
}

func (r *RepeatedField) IsPrimitive() bool {
	return r.Primitive
}

type BasicField struct {
	Optional  bool
	Type      string
	Name      string
	Primitive bool
}

func (r *BasicField) GetType() string {
	return r.Type
}

func (b *BasicField) ToProto(idx int) (string, int) {
	str := fmt.Sprintf("%s %s = %d;", b.Type, b.Name, idx)
	if b.Optional {
		str = "optional " + str
	}
	return str, idx
}

func (r *BasicField) IsPrimitive() bool {
	return r.Primitive
}

type OneOfField struct {
	Name      string
	Types     []*BasicField
	Primitive bool
}

func (r *OneOfField) GetType() string {
	panic("OneOfField does not have a type")
}

func (o *OneOfField) ToProto(idx int) (string, int) {
	var sb strings.Builder

	sb.WriteString("oneof " + o.Name + " {\n")
	for _, field := range o.Types {
		sb.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", field.Type, field.Name, idx))
		idx++
	}
	sb.WriteString("\t}")
	idx--
	return sb.String(), idx
}

func (r *OneOfField) IsPrimitive() bool {
	return r.Primitive
}
