package utils

import (
	"regexp"
	"strings"

	"github.com/gobeam/stringy"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchUnderscoreNum = regexp.MustCompile(`(_)(\d)`)

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchUnderscoreNum.ReplaceAllString(snake, "${2}")
	return strings.ToLower(snake)
}

var pattern1 = regexp.MustCompile(`(\d)([a-z])`)
var pattern2 = regexp.MustCompile(`([a-z])(\d)`)
var underscoreBeforeNum = regexp.MustCompile(`(_)(\d)`)

func ToPascalCase(str string, modifier ...func(string) string) string {
	result := stringy.New(str).PascalCase().Get()

	for _, m := range modifier {
		result = m(result)
	}

	result = DirtyFixes(result)

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
func RemoveUnderscoreBeforeNum(input string) string {
	// Use the ReplaceAllStringFunc method to apply a function to each match
	return underscoreBeforeNum.ReplaceAllStringFunc(input, func(match string) string {
		return strings.ToUpper(string(match[1]))
	})
}

func DirtyFixes(input string) string {
	return strings.Replace(input, "VRF", "Vrf", -1)
}
