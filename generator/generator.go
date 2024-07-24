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

	seenFields map[string]bool
	IdToField  map[int64]protobuf.Field
}

func NewGenerator(templatePath string, messages []*protobuf.Message, metadata *types.Metadata, IdToField map[int64]protobuf.Field) *Generator {
	return &Generator{
		templatePath: templatePath,
		Messages:     messages,
		Metadata:     metadata,
		seenFields:   make(map[string]bool),
		IdToField:    IdToField,
	}
}

func (g *Generator) Generate() error {
	b, err := os.ReadFile(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	funcMap := template.FuncMap{
		"wrap":        g.Wrap,
		"newWrapItem": NewWrapItems,
	}

	templ, err := template.New("").Funcs(funcMap).Parse(string(b))
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

func (g *Generator) IsSeen(k string) bool {
	return g.seenFields[k]
}

func (g *Generator) Seen(k string) bool {
	g.seenFields[k] = true
	return true
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

func (g *Generator) IsOneOf(field protobuf.Field) bool {
	if _, ok := field.(*protobuf.OneOfField); ok {
		return true
	}
	return false
}

type WrapItem struct {
	Key   any
	Value any
}

func NewWrapItems(a ...any) []WrapItem {
	items := make([]WrapItem, len(a))
	for i := 0; i < len(a); i += 2 {
		items[i] = WrapItem{
			Key:   a[i],
			Value: a[i+1],
		}

	}

	return items
}
func (g *Generator) VariantChildIds(id int64) []string {
	var out []string
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[id]
	variant := ttype.Def.Variant
	for _, v := range variant.Variants {
		var o []int64 //child of 94
		for _, f := range v.Fields {
			o = append(o, f.Type.Int64())
		}
		s := "[]int64{"
		for _, v := range o {
			s += fmt.Sprintf("%v, ", v)
		}
		s = s[:len(s)-2]
		s += "}"

		out = append(out, s)
	}

	return out
}

func (g *Generator) Wrap(items []WrapItem) map[any]any {
	out := make(map[any]any)
	for _, item := range items {
		out[item.Key] = item.Value
	}
	out["generator"] = g
	return out
}
