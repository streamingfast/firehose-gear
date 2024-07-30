package generator

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
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
		"hasSuffix":   strings.HasSuffix,
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

func (g *Generator) Wrap(items []WrapItem) map[any]any {
	out := make(map[any]any)
	for _, item := range items {
		out[item.Key] = item.Value
	}
	out["generator"] = g
	return out
}

func (g *Generator) TypeForField(field protobuf.Lookupable) string {

	if field.GetName() == "sp_arithmetic_per_things_PerU16" {
		fmt.Println("")
	}

	idx := field.GetLookupId()
	if idx == math.MaxInt64 {
		return "registry.DecodedFields"
	}
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]

	switch {
	case ttype.Def.IsCompact:
		//return "types.UCompact"
	case ttype.Def.IsPrimitive:
		switch ttype.Def.Primitive.Si0TypeDefPrimitive {
		case types.IsU8:
			return "types.U8"
		case types.IsU16:
			return "types.U16"
		case types.IsU32:
			return "types.U32"
		case types.IsU64:
			return "types.U64"
		case types.IsU128:
			return "types.U128"
		case types.IsU256:
			return "types.U256"
		case types.IsI8:
			return "types.I8"
		case types.IsI16:
			return "types.I16"
		case types.IsI32:
			return "types.I32"
		case types.IsI64:
			return "types.I64"
		case types.IsI128:
			return "types.I128"
		case types.IsI256:
			return "types.I256"
		case types.IsChar:
			panic("Not implemented")
		case types.IsStr:
			return "types.Text"
		case types.IsBool:
			return "bool"
		}

		//case ttype.Def.IsSequence:
		//	return "*types.Si1TypeDefSequence"
		//case ttype.Def.IsArray:
		//	return "*types.Si1TypeDefArray"
		//case ttype.Def.IsTuple:
		//return "*types.Si1TypeDefTuple"
		//case ttype.Def.IsVariant:
		//	types.Si1Variant{}
		//	return "*types.Si1TypeDefTuple"
		//case ttype.Def.IsComposite:
	}
	return "registry.DecodedFields"
}

func (g *Generator) CompactChildType(field protobuf.Field) types.Si1TypeDef {
	idx := field.GetLookupId()
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]
	return ttype.Def
}

func (g *Generator) LookupType(idx int64) types.Si1TypeDef {
	if idx == math.MaxInt64 {
		return types.Si1TypeDef{}
	}
	ttype := g.Metadata.AsMetadataV14.EfficientLookup[idx]
	return ttype.Def
}

func (g *Generator) FuncNameForMessage(msg *protobuf.Message) string {
	return "To_" + msg.FullTypeName()
}
func (g *Generator) FuncNameForOneOf(msg *protobuf.Message, field protobuf.Field) string {
	return "Set_OneOf_" + msg.FullTypeName() + "_" + field.GetName()
}

func (g *Generator) FuncNameForOptional(msg *protobuf.Message, field protobuf.Field) string {
	return "To_Optional_" + msg.FullTypeName() + "_" + field.GetName()
}

func (g *Generator) FuncNameForRepeated(msg *protobuf.Message, field protobuf.Field) string {
	return "To_Repeated_" + msg.FullTypeName() + "_" + field.GetName()
}

func (g *Generator) FuncNameForField(msg *protobuf.Message, field protobuf.Field) string {
	return "To_" + msg.FullTypeName() + "_" + field.GetName()
}

func (g *Generator) FuncNameForPrimitive(msg *protobuf.Message, field protobuf.Field) string {
	if field.IsOptional() {
		return "To_Optional_" + field.GetType()
	}

	if field.IsRepeated() {
		return "To_Repeated_" + field.GetType()
	}

	return "To_" + field.GetType()
}

func (g *Generator) ReturnTypeForOneOf(msg *protobuf.Message, field protobuf.Field) string {
	return "pbgear." + msg.FullTypeName() + "_" + field.GetType()
}
