package utils

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func Formatter(s string) string {
	normStr := norm.NFD.String(s)
	var result []rune

	for _, r := range normStr {
		if !unicode.Is(unicode.Mn, r) {
			result = append(result, r)
		}
	}
	return string(result)
}
