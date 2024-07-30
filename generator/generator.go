package generator

import (
	"bytes"
	"fmt"
	"math"
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

func NewGenerator(templatePath string, messages []*protobuf.Message, metadata *types.Metadata) *Generator {
	return &Generator{
		templatePath: templatePath,
		Messages:     messages,
		Metadata:     metadata,
		seenFields:   make(map[string]bool),
	}
}

func (g *Generator) Generate() ([]byte, error) {
	b, err := os.ReadFile(g.templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	funcMap := template.FuncMap{
		"wrap":        g.Wrap,
		"newWrapItem": NewWrapItems,
	}

	templ, err := template.New("").Funcs(funcMap).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	buffer := &bytes.Buffer{}
	if err := templ.Execute(buffer, g); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buffer.Bytes(), nil
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
		if len(v.Fields) == 0 {
			return []string{"[]int64{}"}
		}
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

func (g *Generator) TypeForField(field protobuf.Lookupable) string {
	idx := field.GetLookupId()
	if idx == math.MaxInt64 {
		return "registry.DecodedFields"
	}
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]

	switch {
	case ttype.Def.IsCompact:
		return "types.UCompact"
		//case ttype.Def.IsSequence:
		//	return "*types.Si1TypeDefSequence"
		//case ttype.Def.IsArray:
		//	return "*types.Si1TypeDefArray"
		//case ttype.Def.IsTuple:
		//return "*types.Si1TypeDefTuple"
		//case ttype.Def.IsVariant:
		//	types.Si1Variant{}
		//	return "*types.Si1TypeDefTuple"
		//case ttype.Def.IsCompact:

		//case ttype.Def.IsComposite:
	}
	return "registry.DecodedFields"
}

func (g *Generator) CompactChildType(field protobuf.Field) types.Si1TypeDef {
	idx := field.GetLookupId()
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]
	return ttype.Def
}

//func (g *Generator) CompactToValue(field protobuf.Field) string {
//	idx := field.GetLookupId()
//	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]
//	childType := g.Metadata.AsMetadataV14.EfficientLookup[]
//
//	switch {
//	case childType.Def.:
//
//	}
//
//	return
//}
