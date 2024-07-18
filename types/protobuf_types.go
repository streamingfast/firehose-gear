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

	sb.WriteString("message Params {\n")
	for i, message := range p.Messages {
		sb.WriteString(fmt.Sprintf("%s = %d;\n", message.ToOneOfCallFunctions(), i+1))
	}
	sb.WriteString("}\n")

	for _, message := range p.Messages {
		sb.WriteString(message.ToProtoMessage())
	}

	return sb.String()
}

type ProtobufMessage struct {
	Name string
}

func (p *ProtobufMessage) ToProtoMessage(options ...string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nmessage %s {\n", p.Name))
	sb.WriteString("}\n")
	return sb.String()
}

func (p *ProtobufMessage) ToOneOfCallFunctions() string {
	str := stringy.New(p.Name)
	return fmt.Sprintf("\t\t%s %s", str.PascalCase().Get(), utils.ToSnakeCase(p.Name))
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
