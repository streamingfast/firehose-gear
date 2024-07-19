package protobuf

import (
	"fmt"
	"strings"
)

type Field interface {
	SetName(name string)
	ToProto() string
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
	Name   string // Utility_Batch_Call, Utility_Batch_Storage
	Fields []Field
}

func (m *Message) ToProto() string {
	var sb strings.Builder

	sb.WriteString("message " + m.Name + " {\n")
	for _, field := range m.Fields {
		sb.WriteString("\t" + field.ToProto() + "\n")
	}
	sb.WriteString("}\n")

	return sb.String()
}

type RepeatedField struct {
	Type string
	Name string
}

func (r *RepeatedField) SetName(name string) {
	r.Name = name
}

func (r *RepeatedField) ToProto() string {
	return "repeated " + r.Type + " " + r.Name
}

type BasicField struct {
	Optional bool
	Type     string
	Name     string
}

func (b *BasicField) SetName(name string) {
	b.Name = name
}

func (b *BasicField) ToProto() string {
	if b.Optional {
		return "optional " + b.Type + " " + b.Name
	}
	return b.Type + " " + b.Name
}

type OneOfField struct {
	Name  string
	Types []string
}

func (o *OneOfField) SetName(name string) {
	o.Name = name
}

func (o *OneOfField) ToProto() string {
	var sb strings.Builder

	sb.WriteString("oneof " + o.Name + " {\n")
	for i, field := range o.Types {
		sb.WriteString(fmt.Sprintf("\t%s = %d;\n", field, i+1))
	}
	sb.WriteString("}")

	return sb.String()
}
