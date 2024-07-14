package types

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/streamingfast/firehose-gear/utils"
)

type ProtobufMessages struct {
	SyntaxName  string // default: proto3
	PackageName string // default: sf.gear.extrinsics.type.v1
	Messages    []*ProtobufMessage
}

func (p *ProtobufMessages) GetSyntaxName() string {
	return fmt.Sprintf("syntax = \"%s\";\n", p.SyntaxName)
}

func (p *ProtobufMessages) GetPackageName() string {
	return fmt.Sprintf("package %s;\n\n", p.PackageName)
}

func (p *ProtobufMessages) ToProtoMessage() string {
	var sb strings.Builder
	sb.WriteString(p.GetSyntaxName())
	sb.WriteString(p.GetPackageName())

	for _, message := range p.Messages {
		sb.WriteString(message.ToProtoMessage())
	}

	return sb.String()
}

// All the fields that can be present in a protobuf
type ProtobufMessage struct {
	Name          string
	CallFunctions *CallFunctions
	Messages      []*ProtobufMessage
	ProtobufTypes []*ProtobufType
}

func (p *ProtobufMessage) ToProtoMessage(options ...string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nmessage %s {\n", p.Name))

	if p.CallFunctions != nil {
		sb.WriteString(p.CallFunctions.ToProtoMessage())
	}

	for i, pt := range p.ProtobufTypes {
		sb.WriteString(fmt.Sprintf("\t%s = %d;\n", pt.ToProtoMessage(), i+1))
	}

	for _, m := range p.Messages {
		sb.WriteString(fmt.Sprintf("\t%s", m.ToProtoMessage()))
	}

	sb.WriteString("}\n")
	return sb.String()
}

type CallFunctionProtobufMessageType struct {
	MessageType string
}

func (p *CallFunctionProtobufMessageType) ToProtoMessage() string {
	messageTypeStr := stringy.New(p.MessageType)
	pascalStr := messageTypeStr.PascalCase().Get()
	snakeStr := utils.ToSnakeCase(p.MessageType)
	return fmt.Sprintf("%s %s", pascalStr, snakeStr)
}

// Key value pair of Name and Type of the protobuf type
type ProtobufType struct {
	Name string
	Type string
}

func (p *ProtobufType) ToProtoMessage() string {
	return fmt.Sprintf("%s %s", p.Type, p.Name)
}
