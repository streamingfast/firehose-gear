package types

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/streamingfast/firehose-gear/utils"
)

type IType interface {
	ToProtoMessage(options ...string) string
	GetName() string
}

type Types struct {
	Types []IType
}

type BaseType struct {
	Optional bool
}

type Composite struct {
	BaseType
	CompositeFields []*CompositeField
}

func (c *Composite) GetName() string {
	return "Composite"
}

func (c *Composite) ToProtoMessage(options ...string) string {
	var sb strings.Builder
	for _, f := range c.CompositeFields {
		sb.WriteString(f.Type.ToProtoMessage())
	}
	return sb.String()
}

type CompositeField struct {
	Name string
	Type IType
}

func (f *CompositeField) ToProtoMessage(options ...string) string {
	if f.Name == "" {
		return f.Type.ToProtoMessage()
	}

	return fmt.Sprintf("%s %s", f.Type.ToProtoMessage(), f.Name)
}

func (f *CompositeField) GetName() string {
	str := stringy.New(f.Name)
	return str.PascalCase().Get()
}

type Variants struct {
	BaseType
	Variants []*Variant
}

func (v *Variants) GetName() string {
	return "Variants"
}

func (v *Variants) VariantNames() string {
	var sb strings.Builder

	for _, variant := range v.Variants {
		str := stringy.New(variant.Name)
		pascalStr := str.PascalCase().Get()
		snakeStr := utils.ToSnakeCase(variant.Name)
		sb.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", pascalStr, snakeStr, variant.Index))
	}
	return sb.String()
}

func (v *Variants) ToProtoMessage(options ...string) string {
	var sb strings.Builder

	for _, variant := range v.Variants {
		sb.WriteString(variant.ToProtoMessage())
	}

	return sb.String()
}

type Variant struct {
	Name          string
	VariantFields []*VariantField
	Index         uint64
}

func (v *Variant) GetName() string {
	str := stringy.New(v.Name)
	return fmt.Sprintf("Variant%s", str.PascalCase().Get())
}

func (v *Variant) ToProtoMessage(options ...string) string {
	var sb strings.Builder

	tabs := "\t"
	str := stringy.New(v.Name)
	sb.WriteString(fmt.Sprintf("%smessage %s {\n", tabs, str.PascalCase().Get()))
	for i, field := range v.VariantFields {
		sb.WriteString(fmt.Sprintf("\t\t%s = %d;\n", field.ToProtoMessage(), i))
	}
	sb.WriteString(fmt.Sprintf("%s}\n", tabs))
	return sb.String()
}

type VariantField struct {
	Name string
	Type IType
}

func (f *VariantField) ToProtoMessage(options ...string) string {
	return fmt.Sprintf("%s %s", f.Type.ToProtoMessage(), f.Name)
}

func (f *VariantField) GetName() string {
	str := stringy.New(f.Name)
	return fmt.Sprintf("VariantField%s", str.PascalCase().Get())
}

func (f *VariantField) ToProtobufType() *ProtobufType {
	return &ProtobufType{
		Name: f.Name,
		Type: f.Type.ToProtoMessage(),
	}
}

type Sequence struct {
	Item IType
}

func (s *Sequence) GetName() string {
	return fmt.Sprintf("Sequence%s", s.Item.GetName())
}

func (s *Sequence) ToProtoMessage(options ...string) string {
	t, ok := s.Item.(*Primitive)
	if ok {
		if t.Si0TypeDefPrimitive.IsU8 {
			return "bytes"
		}
	}

	tuple, ok := s.Item.(*Tuple)
	if ok {
		return tuple.ToProtoMessage()
	}

	seq, ok := s.Item.(*Sequence)
	if ok {
		return fmt.Sprintf("repeated %s", seq.ToProtoMessage())
	}

	return fmt.Sprintf("repeated %s\n", s.Item.ToProtoMessage())
}

type Array struct {
	Type IType
}

func (a *Array) GetName() string {
	return fmt.Sprintf("Array%s", a.Type.GetName())
}

func (a *Array) ToProtoMessage(options ...string) string {
	t, ok := a.Type.(*Primitive)
	if ok {
		if t.Si0TypeDefPrimitive.IsU8 {
			return "bytes"
		}
	}

	var sb strings.Builder
	sb.WriteString(a.Type.ToProtoMessage())
	return sb.String()
}

type Tuple struct {
	Name  string
	Items []IType
}

func (t *Tuple) GetName() string {
	str := stringy.New(t.Name)
	return fmt.Sprintf("Tuple%s", str.PascalCase().Get())
}

func (t *Tuple) ToProtoMessage(options ...string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\n\t\tmessage Tuple_%s {\n", t.Name))
	for i, item := range t.Items {
		sb.WriteString(fmt.Sprintf("\t\t\t%s %s_%d", item.ToProtoMessage(), t.Name, i+1))
		sb.WriteString(fmt.Sprintf(" = %d;\n", i+1))
	}
	sb.WriteString("\t\t}\n")
	return sb.String()
}

type Primitive struct {
	Si0TypeDefPrimitive *Type
}

func (p *Primitive) GetName() string {
	return p.Si0TypeDefPrimitive.ToProtoMessage()
}

func (p *Primitive) GetProtoFieldName() string {
	return p.Si0TypeDefPrimitive.ToProtoMessage()
}

func (p *Primitive) ToProtoMessage(options ...string) string {
	return p.Si0TypeDefPrimitive.ToProtoMessage()
}

type Compact struct {
	Type IType
}

func (c *Compact) GetName() string {
	return fmt.Sprintf("Compact%s", c.Type.GetName())
}

func (c *Compact) ToProtoMessage(options ...string) string {
	return c.Type.ToProtoMessage()
}

type BitSequence struct {
	BitStoreType IType
	BitOrderType IType
}

func (b *BitSequence) GetName() string {
	return "BitSequence"
}

func (b *BitSequence) ToProtoMessage(options ...string) string {
	panic("not implemented")
}

type HistoricMetaCompat string

func (h HistoricMetaCompat) GetName() string {
	return "HistoricMetaCompat"
}

func (h HistoricMetaCompat) ToProtoMessage(options ...string) string {
	panic("not implemented")
}

type Type struct {
	IsBool bool
	IsChar bool
	IsStr  bool
	IsU8   bool
	IsU16  bool
	IsU32  bool
	IsU64  bool
	IsU128 bool
	IsU256 bool
	IsI8   bool
	IsI16  bool
	IsI32  bool
	IsI64  bool
	IsI128 bool
	IsI256 bool
}

func (t *Type) GetProtoFieldName() string {
	if t.IsBool {
		return "Bool"
	}

	if t.IsChar {
		return "Char"
	}

	if t.IsStr {
		return "String"
	}

	if t.IsU128 {
		return "Uint128"
	}

	if t.IsU256 {
		return "Uint256"
	}

	if t.IsI128 {
		return "Int128"
	}

	if t.IsI256 {
		return "Int256"
	}

	if t.IsU8 {
		return "Uint8"
	}

	if t.IsU16 {
		return "Uint16"
	}

	if t.IsU32 {
		return "Uint32"
	}

	if t.IsU64 {
		return "Uint64"
	}

	if t.IsI8 {
		return "Int8"
	}

	if t.IsI16 {
		return "Int16"
	}

	if t.IsI32 {
		return "Int32"
	}

	if t.IsI64 {
		return "Int64"
	}

	panic("unknown type")
}

func (t *Type) ToProtoMessage(options ...string) string {
	if t.IsBool {
		return "bool"
	}

	if t.IsChar || t.IsStr || t.IsU128 || t.IsU256 || t.IsI128 || t.IsI256 {
		return "string"
	}

	if t.IsU8 || t.IsU16 || t.IsU32 {
		return "uint32"
	}

	if t.IsU64 {
		return "uint64"
	}

	if t.IsI8 || t.IsI16 || t.IsI32 {
		return "int32"
	}

	if t.IsI64 {
		return "int64"
	}

	return ""
}
