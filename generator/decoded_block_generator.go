package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/streamingfast/firehose-gear/protobuf"
)

type DecodedBlockGenerator struct {
	templatePath string
	Messages     []*protobuf.Message
}

func NewDecodedBlockGenerator(templatePath string, messages []*protobuf.Message) *DecodedBlockGenerator {
	return &DecodedBlockGenerator{
		templatePath: templatePath,
		Messages:     messages,
	}
}

func (g *DecodedBlockGenerator) Generate() ([]byte, error) {
	b, err := os.ReadFile(g.templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
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
