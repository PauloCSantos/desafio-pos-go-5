package services

import (
	"unicode"
)

func ValidateCep(cep string) (valid bool) {
	if len(cep) != 8 {
		return false
	}
	for _, char := range cep {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
