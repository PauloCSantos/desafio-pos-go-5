package services

import "unicode"

func ValidateInput(input string) bool {
	return len(input) == 8 && isString(input)
}

func isString(input string) bool {
	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
