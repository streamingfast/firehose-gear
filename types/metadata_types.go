package types

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/streamingfast/firehose-gear/utils"
)

type IType interface {
	ToProtoMessage() string
}

type Types struct {
	Types []IType
}

type CallFunctions struct {
	MessageTypes []*CallFunctionProtobufMessageType
	Variant      *Variants
}

func (c *CallFunctions) ToProtoMessage(options ...string) string {
	var sb strings.Builder
	sb.WriteString("\toneof CallFunctions {\n")

	for i, mt := range c.MessageTypes {
		sb.WriteString(fmt.Sprintf("\t\t%s = %d;\n", mt.ToProtoMessage(), i+1))
	}

	sb.WriteString("\t}\n")
	return sb.String()
}

type Composite struct {
	CompositeFields []*CompositeField
}

func (c *Composite) ToProtoMessage() string {
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

func (f *CompositeField) ToProtoMessage() string {
	if f.Name == "" {
		return f.Type.ToProtoMessage()
	}

	return fmt.Sprintf("%s %s", f.Type.ToProtoMessage(), f.Name)
}

type VariantField struct {
	Name string
	Type IType
}

func (f *VariantField) ToProtoMessage() string {
	return fmt.Sprintf("%s %s", f.Type.ToProtoMessage(), f.Name)
}

func (f *VariantField) ToProtobufType() *ProtobufType {
	return &ProtobufType{
		Name: f.Name,
		Type: f.Type.ToProtoMessage(),
	}
}

type Variants struct {
	Variants []*Variant
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

func (v *Variants) ToCallFunctions() *CallFunctions {
	var cfMessageTypes []*CallFunctionProtobufMessageType
	for _, variant := range v.Variants {
		// if a variant is a simple type, then we don't need to create a call function for it
		if variant.isSimpleType() {
			continue
		}

		cfMessageTypes = append(cfMessageTypes, &CallFunctionProtobufMessageType{
			MessageType: variant.Name,
		})
	}
	return &CallFunctions{
		MessageTypes: cfMessageTypes,
		Variant:      v,
	}
}

func (v *Variants) ToProtobufType() []*ProtobufType {
	var protobufTypes []*ProtobufType
	for _, variant := range v.Variants {
		if variant.isSimpleType() {
			protobufTypes = append(protobufTypes, variant.ToProtobufType()...)
		}
	}
	return protobufTypes
}

func (v *Variants) ToProtobufMessages() []*ProtobufMessage {
	var protobufMessages []*ProtobufMessage
	for _, variant := range v.Variants {
		protobufMessages = append(protobufMessages, variant.ToProtobufMessage())
	}
	return protobufMessages
}

func (v *Variants) ToProtoMessage() string {
	var sb strings.Builder

	if v.isAnOption() {
		sb.WriteString(fmt.Sprintf("optional %s", v.Variants[1].VariantFields[0].Type.ToProtoMessage()))
		return sb.String()
	}

	for _, variant := range v.Variants {
		sb.WriteString(variant.ToProtoMessage())
	}

	return sb.String()
}

func (v *Variants) isAnOption() bool {
	isNonePresent := false
	isSomePresent := false
	for _, vari := range v.Variants {
		if vari.Name == "None" {
			isNonePresent = true
			continue
		}

		if vari.Name == "Some" {
			isSomePresent = true
			continue
		}
	}
	return isNonePresent && isSomePresent
}

type Variant struct {
	Name          string
	VariantFields []*VariantField
	Index         uint64
}

func (v *Variant) ToProtoMessage() string {
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

func (v *Variant) ToProtobufMessage() *ProtobufMessage {
	var protobufTypes []*ProtobufType
	for _, field := range v.VariantFields {
		protobufTypes = append(protobufTypes, &ProtobufType{
			Name: v.Name,
			Type: field.Type.ToProtoMessage(),
		})
	}

	return &ProtobufMessage{
		Name:          v.Name,
		ProtobufTypes: protobufTypes,
	}
}

func (v *Variant) isSimpleType() bool {
	for _, field := range v.VariantFields {
		if field.Name != "" {
			return false
		}
	}
	return true
}

func (v *Variant) ToProtobufType() []*ProtobufType {
	var protobufTypes []*ProtobufType
	if v.isSimpleType() {
		for _, field := range v.VariantFields {
			protobufTypes = append(protobufTypes, field.ToProtobufType())
		}
	}
	return protobufTypes
}

type Sequence struct {
	Item IType
}

// TODO: fix this
func (s *Sequence) ToProtoMessage() string {
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

// TODO: fix this
func (a *Array) ToProtoMessage() string {
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

func (t *Tuple) ToProtoMessage() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\n\t\tmessage Tuple_%s {\n", t.Name))
	for i, item := range t.Items {
		sb.WriteString(fmt.Sprintf("\t\t\t%s %s_%d", item.ToProtoMessage(), t.Name, i+1))
		sb.WriteString(fmt.Sprintf(" = %d;\n", i+1))
	}
	sb.WriteString("\t\t}\n")
	return sb.String()
}

func (t *Tuple) ToProtobufMessage() *ProtobufMessage {
	var protobufTypes []*ProtobufType

	for i, item := range t.Items {
		protobufTypes = append(protobufTypes, &ProtobufType{
			Name: fmt.Sprintf("%s_%d", t.Name, i+1),
			Type: item.ToProtoMessage(),
		})
	}

	return &ProtobufMessage{
		Name:          fmt.Sprintf("Tuple_%s", t.Name),
		ProtobufTypes: protobufTypes,
	}
}

type Primitive struct {
	Si0TypeDefPrimitive *Type
}

func (p *Primitive) ToProtoMessage() string {
	return p.Si0TypeDefPrimitive.ToProtoMessage()
}

type Compact struct {
	Type IType
}

func (c *Compact) ToProtoMessage() string {
	return c.Type.ToProtoMessage()
}

type BitSequence struct {
	BitStoreType IType
	BitOrderType IType
}

func (b *BitSequence) ToProtoMessage() string {
	panic("not implemented")
}

type HistoricMetaCompat string

func (h HistoricMetaCompat) ToProtoMessage() string {
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
