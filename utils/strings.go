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

var pattern1 = regexp.MustCompile(`(\d)([a-z])`)
var pattern2 = regexp.MustCompile(`([a-z])(\d)`)

func ToPascalCase(str string, modifier ...func(string) string) string {
	result := stringy.New(str).PascalCase().Get()

	for _, m := range modifier {
		result = m(result)
	}

	return result
}

func CapitalizeCharAfterNum(input string) string {
	// Use the ReplaceAllStringFunc method to apply a function to each match
	return pattern1.ReplaceAllStringFunc(input, func(match string) string {
		return string(match[0]) + strings.ToUpper(string(match[1]))
	})
}
func UnderscoreBetweenLetterAndNum(input string) string {
	// Use the ReplaceAllStringFunc method to apply a function to each match
	return pattern2.ReplaceAllStringFunc(input, func(match string) string {
		return string(match[0]) + "_" + string(match[1])
	})
}
