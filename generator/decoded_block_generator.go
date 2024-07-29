package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type DecodedBlockGenerator struct {
	templatePath string
}

func NewDecodedBlockGenerator(templatePath string) *DecodedBlockGenerator {
	return &DecodedBlockGenerator{
		templatePath: templatePath,
	}
}

func (g *DecodedBlockGenerator) Generate() error {
	b, err := os.ReadFile(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	templ, err := template.New("").Funcs(nil).Parse(string(b))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	buffer := &bytes.Buffer{}
	if err := templ.Execute(buffer, g); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	content := buffer.Bytes()
	err = os.WriteFile("../templates/decoded_block.proto", content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
