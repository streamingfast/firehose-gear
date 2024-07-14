package types

import (
	"fmt"
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

func Test_RuntimeModule_Gear_CallFunctions(t *testing.T) {
	p := &ProtobufMessages{
		SyntaxName:  "proto3",
		PackageName: "sf.gear.extrinsics.type.v1",
		Messages: []*ProtobufMessage{
			{
				Name: "Gear",
				CallFunctions: &CallFunctions{
					Variant: &Variants{
						Variants: []*Variant{
							{
								Name: "upload_code",
								VariantFields: []*VariantField{
									{
										Name: "code",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
								},
								Index: 1,
							},
							{
								Name: "upload_program",
								VariantFields: []*VariantField{
									{
										Name: "code",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
									{
										Name: "salt",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
									{
										Name: "init_payload",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
									{
										Name: "gas_limit",
										Type: &Primitive{
											Si0TypeDefPrimitive: &Type{
												IsU64: true,
											},
										},
									},
									{
										Name: "value",
										Type: &Primitive{
											Si0TypeDefPrimitive: &Type{
												IsU128: true,
											},
										},
									},
									{
										Name: "keep_alive",
										Type: &Primitive{
											Si0TypeDefPrimitive: &Type{
												IsBool: true,
											},
										},
									},
								},
								Index: 2,
							},
							// {
							// 	Name: "create_program",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "code_id",
							// 			Type: &Composite{
							// 				Fields: []*Field{
							// 					{
							// 						Type: &Array{
							// 							Type: &Primitive{
							// 								Si0TypeDefPrimitive: &Type{
							// 									IsU8: true,
							// 								},
							// 							},
							// 						},
							// 					},
							// 				},
							// 			},
							// 		},
							// {
							// 	Name: "salt",
							// 	Type: &Sequence{
							// 		Item: &Primitive{
							// 			Si0TypeDefPrimitive: &Type{
							// 				IsU8: true,
							// 			},
							// 		},
							// 	},
							// },
							// {
							// 	Name: "init_payload",
							// 	Type: &Sequence{
							// 		Item: &Primitive{
							// 			Si0TypeDefPrimitive: &Type{
							// 				IsU8: true,
							// 			},
							// 		},
							// 	},
							// },
							// {
							// 	Name: "gas_limit",
							// 	Type: &Primitive{
							// 		Si0TypeDefPrimitive: &Type{
							// 			IsU64: true,
							// 		},
							// 	},
							// },
							// {
							// 	Name: "value",
							// 	Type: &Primitive{
							// 		Si0TypeDefPrimitive: &Type{
							// 			IsU128: true,
							// 		},
							// 	},
							// },
							// {
							// 	Name: "keep_alive",
							// 	Type: &Primitive{
							// 		Si0TypeDefPrimitive: &Type{
							// 			IsBool: true,
							// 		},
							// 	},
							// },
							// 	},
							// 	Index: 3,
							// },
							// {
							// 	Name: "send_message",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "destination",
							// 			Type: &Composite{
							// 				Fields: []*Field{
							// 					{
							// 						Type: &Primitive{
							// 							Si0TypeDefPrimitive: &Type{
							// 								IsU8: true,
							// 							},
							// 						},
							// 					},
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "payload",
							// 			Type: &Sequence{
							// 				Item: &Primitive{
							// 					Si0TypeDefPrimitive: &Type{
							// 						IsU8: true,
							// 					},
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "gas_limit",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsU64: true,
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "value",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsU128: true,
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "keep_alive",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsBool: true,
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 4,
							// },
							// {
							// 	Name: "send_reply",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "reply_to_id",
							// 			Type: &Composite{
							// 				Fields: []*Field{
							// 					{
							// 						Type: &Primitive{
							// 							Si0TypeDefPrimitive: &Type{
							// 								IsU8: true,
							// 							},
							// 						},
							// 					},
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "payload",
							// 			Type: &Sequence{
							// 				Item: &Primitive{
							// 					Si0TypeDefPrimitive: &Type{
							// 						IsU8: true,
							// 					},
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "gas_limit",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsU64: true,
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "value",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsU128: true,
							// 				},
							// 			},
							// 		},
							// 		{
							// 			Name: "keep_alive",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsBool: true,
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 5,
							// },
							// {
							// 	Name: "claim_value",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "message_id",
							// 			Type: &Composite{
							// 				Fields: []*Field{
							// 					{
							// 						Type: &Primitive{
							// 							Si0TypeDefPrimitive: &Type{
							// 								IsU8: true,
							// 							},
							// 						},
							// 					},
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 6,
							// },
							// {
							// 	Name: "run",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "max_gas",
							// 			Type: &Variants{
							// 				Variants: []*Variant{
							// 					{
							// 						Name: "None",
							// 					},
							// 					{
							// 						Name: "Some",
							// 						Fields: []*Field{
							// 							{
							// 								Name: "",
							// 								Type: &Primitive{
							// 									Si0TypeDefPrimitive: &Type{
							// 										IsU64: true,
							// 									},
							// 								},
							// 							},
							// 						},
							// 					},
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 7,
							// },
							// {
							// 	Name: "set_execute_inherent",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "message_id",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsBool: true,
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 8,
							// },
						},
					},
				},
			},
		},
	}

	fmt.Println(p.ToProtoMessage())
	require.True(t, false)
}

func Test_RuntimeModule_System_CallFunctions(t *testing.T) {
	p := &ProtobufMessages{
		SyntaxName:  "proto3",
		PackageName: "sf.gear.extrinsics.type.v1",
		Messages: []*ProtobufMessage{
			{
				Name: "System",
				CallFunctions: &CallFunctions{
					Variant: &Variants{
						Variants: []*Variant{
							{
								Name: "remark",
								VariantFields: []*VariantField{
									{
										Name: "remark",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
								},
								Index: 1,
							},
							{
								Name: "set_heap_pages",
								VariantFields: []*VariantField{
									{
										Name: "pages",
										Type: &Primitive{
											Si0TypeDefPrimitive: &Type{
												IsU8: true,
											},
										},
									},
								},
								Index: 2,
							},
							{
								Name: "set_code",
								VariantFields: []*VariantField{
									{
										Name: "code",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
								},
								Index: 3,
							},
							{
								Name: "set_code_without_checks",
								VariantFields: []*VariantField{
									{
										Name: "code",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
								},
								Index: 4,
							},
							{
								Name: "set_storage",
								VariantFields: []*VariantField{
									{
										Name: "items",
										Type: &Sequence{
											Item: &Tuple{
												Name: "items",
												Items: []IType{
													&Sequence{
														Item: &Primitive{
															Si0TypeDefPrimitive: &Type{
																IsU8: true,
															},
														},
													},
													&Sequence{
														Item: &Primitive{
															Si0TypeDefPrimitive: &Type{
																IsU8: true,
															},
														},
													},
												},
											},
										},
									},
								},
								Index: 5,
							},
							{
								Name: "kill_storage",
								VariantFields: []*VariantField{
									{
										Name: "keys",
										Type: &Sequence{
											Item: &Sequence{
												Item: &Primitive{
													Si0TypeDefPrimitive: &Type{
														IsU8: true,
													},
												},
											},
										},
									},
								},
								Index: 6,
							},
							{
								Name: "kill_prefix",
								VariantFields: []*VariantField{
									{
										Name: "prefix",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
									{
										Name: "subkeys",
										Type: &Primitive{
											Si0TypeDefPrimitive: &Type{
												IsU32: true,
											},
										},
									},
								},
								Index: 7,
							},
							{
								Name: "remark_with_event",
								VariantFields: []*VariantField{
									{
										Name: "remark",
										Type: &Sequence{
											Item: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU8: true,
												},
											},
										},
									},
								},
								Index: 8,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(p.ToProtoMessage())
	require.True(t, false)
}

func Test_RuntimeModule_Timestamp_CallFunctions(t *testing.T) {
	p := &ProtobufMessages{
		SyntaxName:  "proto3",
		PackageName: "sf.gear.extrinsics.type.v1",
		Messages: []*ProtobufMessage{
			{
				Name: "Timestamp",
				CallFunctions: &CallFunctions{
					Variant: &Variants{
						Variants: []*Variant{
							{
								Name: "set",
								VariantFields: []*VariantField{
									{
										Name: "now",
										Type: &Compact{
											Type: &Primitive{
												Si0TypeDefPrimitive: &Type{
													IsU64: true,
												},
											},
										},
									},
								},
								Index: 1,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(p.ToProtoMessage())
	require.True(t, false)
}

func Test_RuntimeModule_Babe_CallFunctions(t *testing.T) {
	p := &ProtobufMessages{
		SyntaxName:  "proto3",
		PackageName: "sf.gear.extrinsics.type.v1",
		Messages: []*ProtobufMessage{
			{
				Name: "Babe",
				CallFunctions: &CallFunctions{
					Variant: &Variants{
						Variants: []*Variant{
							{
								Name: "report_equivocation",
								VariantFields: []*VariantField{
									{
										Name: "equivocation_proof",
										Type: &Composite{
											CompositeFields: []*CompositeField{
												{
													Name: "offender",
													Type: &Composite{
														CompositeFields: []*CompositeField{
															{
																Type: &Composite{
																	CompositeFields: []*CompositeField{
																		{
																			Type: &Primitive{
																				Si0TypeDefPrimitive: &Type{
																					IsU8: true,
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									// {
									// 	Name: "key_owner_proof",
									// 	Type: &Sequence{
									// 		Item: &Primitive{
									// 			Si0TypeDefPrimitive: &Type{
									// 				IsU8: true,
									// 			},
									// 		},
									// 	},
									// },
								},
								Index: 1,
							},
							// {
							// 	Name: "report_equivocation_unsigned",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "pages",
							// 			Type: &Primitive{
							// 				Si0TypeDefPrimitive: &Type{
							// 					IsU8: true,
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 2,
							// },
							// {
							// 	Name: "plan_config_change",
							// 	Fields: []*Field{
							// 		{
							// 			Name: "code",
							// 			Type: &Sequence{
							// 				Item: &Primitive{
							// 					Si0TypeDefPrimitive: &Type{
							// 						IsU8: true,
							// 					},
							// 				},
							// 			},
							// 		},
							// 	},
							// 	Index: 3,
							// },
						},
					},
				},
			},
		},
	}

	fmt.Println(p.ToProtoMessage())
	require.True(t, false)
}
