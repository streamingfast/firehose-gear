package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Primitives(t *testing.T) {
	// test cases for the primitive types
	tests := []struct {
		name     string
		ttype    *Primitive
		expected string
	}{
		{
			name: "bool",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsBool: true,
				},
			},
			expected: "bool",
		},
		{
			name: "char",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsChar: true,
				},
			},
			expected: "string",
		},
		{
			name: "string",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsStr: true,
				},
			},
			expected: "string",
		},
		{
			name: "u8",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU8: true,
				},
			},
			expected: "uint32",
		},
		{
			name: "u16",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU16: true,
				},
			},
			expected: "uint32",
		},
		{
			name: "u32",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU32: true,
				},
			},
			expected: "uint32",
		},
		{
			name: "u64",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU64: true,
				},
			},
			expected: "uint64",
		},
		{
			name: "i8",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI8: true,
				},
			},
			expected: "int32",
		},
		{
			name: "i16",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI16: true,
				},
			},
			expected: "int32",
		},
		{
			name: "i32",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI32: true,
				},
			},
			expected: "int32",
		},
		{
			name: "i64",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI64: true,
				},
			},
			expected: "int64",
		},
		{
			name: "u128",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU128: true,
				},
			},
			expected: "string",
		},
		{
			name: "u256",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsU256: true,
				},
			},
			expected: "string",
		},
		{
			name: "i128",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI128: true,
				},
			},
			expected: "string",
		},
		{
			name: "i256",
			ttype: &Primitive{
				Si0TypeDefPrimitive: &Type{
					IsI256: true,
				},
			},
			expected: "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.ttype.ToProtoMessage()
			require.Equal(t, tt.expected, actual)
		})
	}
}
