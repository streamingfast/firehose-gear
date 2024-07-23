package utils

import (
	"regexp"
	"strings"

	"github.com/gobeam/stringy"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToPascalCase(str string) string {
	input := stringy.New(str).PascalCase().Get()
	// Compile the regex pattern to match a number followed by a lowercase letter
	pattern := regexp.MustCompile(`(\d)([a-z])`)

	// Use the ReplaceAllStringFunc method to apply a function to each match
	result := pattern.ReplaceAllStringFunc(input, func(match string) string {
		return string(match[0]) + strings.ToUpper(string(match[1]))
	})

	return result
}
