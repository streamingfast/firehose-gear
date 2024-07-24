package types

import (
	"fmt"

	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func ConvertPrimitiveType(b substrateTypes.Si0TypeDefPrimitive) *Primitive {
	p := &Primitive{}

	switch b {
	case substrateTypes.IsBool:
		p.Si0TypeDefPrimitive = &Type{IsBool: true}
	case substrateTypes.IsU8:
		p.Si0TypeDefPrimitive = &Type{IsU8: true}
	case substrateTypes.IsU16:
		p.Si0TypeDefPrimitive = &Type{IsU16: true}
	case substrateTypes.IsU32:
		p.Si0TypeDefPrimitive = &Type{IsU32: true}
	case substrateTypes.IsU64:
		p.Si0TypeDefPrimitive = &Type{IsU64: true}
	case substrateTypes.IsI8:
		p.Si0TypeDefPrimitive = &Type{IsI8: true}
	case substrateTypes.IsI16:
		p.Si0TypeDefPrimitive = &Type{IsI16: true}
	case substrateTypes.IsI32:
		p.Si0TypeDefPrimitive = &Type{IsI32: true}
	case substrateTypes.IsI64:
		p.Si0TypeDefPrimitive = &Type{IsI64: true}
	case substrateTypes.IsStr:
		p.Si0TypeDefPrimitive = &Type{IsStr: true}
	case substrateTypes.IsChar:
		p.Si0TypeDefPrimitive = &Type{IsChar: true}
	case substrateTypes.IsU128:
		p.Si0TypeDefPrimitive = &Type{IsU128: true}
	case substrateTypes.IsU256:
		p.Si0TypeDefPrimitive = &Type{IsU256: true}
	case substrateTypes.IsI128:
		p.Si0TypeDefPrimitive = &Type{IsI128: true}
	case substrateTypes.IsI256:
		p.Si0TypeDefPrimitive = &Type{IsI256: true}
	default:
		panic(fmt.Sprintf("unsupported primitive type %v", b))
	}

	return p
}

type Primitive struct {
	Si0TypeDefPrimitive *Type
}

func (p *Primitive) GetProtoFieldName() string {
	return p.Si0TypeDefPrimitive.ToGoType()
}

func (p *Primitive) ToGoType(options ...string) string {
	return p.Si0TypeDefPrimitive.ToGoType()
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

func (t *Type) ToGoType(options ...string) string {
	return t.ToProtoType()
}

func (t *Type) ToProtoType(options ...string) string {
	if t.IsBool {
		return "bool"
	}

	if t.IsChar || t.IsStr || t.IsU128 || t.IsU256 || t.IsI128 || t.IsI256 {
		return "string"
	}

	if t.IsU8 {
		return "uint8"
	}

	if t.IsI8 {
		return "int8"
	}

	if t.IsU16 || t.IsU32 {
		return "uint32"
	}

	if t.IsU64 {
		return "uint64"
	}

	if t.IsI16 || t.IsI32 {
		return "int32"
	}

	if t.IsI64 {
		return "int64"
	}

	panic("unknown type")
}
