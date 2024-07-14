package types

import (
	"strings"
	"testing"
)

func Test_ProtobufType(t *testing.T) {
	tests := []struct {
		name         string
		protobufType *ProtobufType
		expected     string
	}{
		{
			name: "uint64 test",
			protobufType: &ProtobufType{
				Name: "code_id",
				Type: "uint64",
			},
			expected: "uint64 code_id",
		},
		{
			name: "string test",
			protobufType: &ProtobufType{
				Name: "title",
				Type: "string",
			},
			expected: "string title",
		},
		{
			name: "bool test",
			protobufType: &ProtobufType{
				Name: "is_active",
				Type: "bool",
			},
			expected: "bool is_active",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.protobufType.ToProtoMessage()
			if actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}

func Test_CallFunctionProtobufMessageType(t *testing.T) {
	tests := []struct {
		name                            string
		callFunctionProtobufMessageType *CallFunctionProtobufMessageType
		expected                        string
	}{
		{
			name: "Gear UploadCode",
			callFunctionProtobufMessageType: &CallFunctionProtobufMessageType{
				MessageType: "UploadCode",
			},
			expected: "UploadCode upload_code",
		},
		{
			name: "Gear UploadProgram",
			callFunctionProtobufMessageType: &CallFunctionProtobufMessageType{
				MessageType: "UploadProgram",
			},
			expected: "UploadProgram upload_program",
		},
		{
			name: "Gear SetExecuteInherent",
			callFunctionProtobufMessageType: &CallFunctionProtobufMessageType{
				MessageType: "SetExecuteInherent",
			},
			expected: "SetExecuteInherent set_execute_inherent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.callFunctionProtobufMessageType.ToProtoMessage()
			if actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}

func Test_ProtobufMessage_With_OnlyCallFunctions(t *testing.T) {
	tests := []struct {
		name            string
		protobufMessage *ProtobufMessage
		expected        string
	}{
		{
			name: "Gear UploadCode only call functions",
			protobufMessage: &ProtobufMessage{
				Name: "Gear",
				CallFunctions: &CallFunctions{
					MessageTypes: []*CallFunctionProtobufMessageType{
						{
							MessageType: "UploadCode",
						},
					},
				},
			},
			expected: `
message Gear {
	oneof CallFunctions {
		UploadCode upload_code = 1;
	}
}
			`,
		},
		{
			name: "Gear UploadProgram",
			protobufMessage: &ProtobufMessage{
				Name: "Gear",
				CallFunctions: &CallFunctions{
					MessageTypes: []*CallFunctionProtobufMessageType{
						{
							MessageType: "UploadProgram",
						},
					},
				},
			},
			expected: `
message Gear {
	oneof CallFunctions {
		UploadProgram upload_program = 1;
	}
}
			`,
		},
		{
			name: "Gear SetExecuteInherent",
			protobufMessage: &ProtobufMessage{
				Name: "Gear",
				CallFunctions: &CallFunctions{
					MessageTypes: []*CallFunctionProtobufMessageType{
						{
							MessageType: "SetExecuteInherent",
						},
					},
				},
			},
			expected: `
message Gear {
	oneof CallFunctions {
		SetExecuteInherent set_execute_inherent = 1;
	}
}
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.protobufMessage.ToProtoMessage()
			if strings.Trim(actual, "\t") != strings.Trim(tt.expected, "\t") {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}

func Test_ProtobufMessage(t *testing.T) {
	tests := []struct {
		name            string
		protobufMessage *ProtobufMessage
		expected        string
	}{
		{
			name: "Gear UploadCode",
			protobufMessage: &ProtobufMessage{
				Name: "Gear",
				CallFunctions: &CallFunctions{
					MessageTypes: []*CallFunctionProtobufMessageType{
						{
							MessageType: "UploadCode",
						},
					},
				},
				Messages: []*ProtobufMessage{
					{
						Name: "UploadCode",
						ProtobufTypes: []*ProtobufType{
							{
								Name: "code",
								Type: "bytes",
							},
						},
					},
				},
			},
			expected: `
message Gear {
	oneof CallFunctions {
		UploadCode upload_code = 1;
	}
	message UploadCode{
		bytes code = 1;
	}			
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.protobufMessage.ToProtoMessage()
			if actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
