package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/streamingfast/firehose-gear/protobuf"
)

type Generator struct {
	templatePath string

	Messages []*protobuf.Message
	Metadata *types.Metadata
}

func NewGenerator(templatePath string, messages []*protobuf.Message, metadata *types.Metadata) *Generator {
	return &Generator{
		templatePath: templatePath,
		Messages:     messages,
		Metadata:     metadata,
	}
}

func (g *Generator) Generate() error {
	b, err := os.ReadFile(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	templ, err := template.New("").Parse(string(b))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	buffer := &bytes.Buffer{}
	if err := templ.Execute(buffer, g); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	content := buffer.Bytes()
	err = os.WriteFile("../templates/gen_types.go", content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (g *Generator) IsRepeatableField(field protobuf.Field) bool {
	if _, ok := field.(*protobuf.RepeatedField); ok {
		return true
	}
	return false
}

func (g *Generator) IsOptionalField(field protobuf.Field) bool {
	if f, ok := field.(*protobuf.BasicField); ok {
		return f.Optional
	}
	return false
}

func (g *Generator) NoneOptionalFieldComparator(field protobuf.Field) string {
	if g.isStringPrimitive(field) {
		return "\"\""
	}

	if g.isBoolPrimitive(field) {
		return "false"
	}

	return "0"
}

func (g *Generator) isStringPrimitive(field protobuf.Field) bool {
	return field.IsPrimitive() && field.GetType() == "string"
}

func (g *Generator) isBoolPrimitive(field protobuf.Field) bool {
	return field.IsPrimitive() && field.GetType() == "bool"
}