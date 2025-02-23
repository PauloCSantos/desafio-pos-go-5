package utils

import "testing"

func TestFormatter(t *testing.T) {
	word := "São Paulo"

	expectedResult := "Sao Paulo"

	result := Formatter(word)

	if result != expectedResult {
		t.Errorf("Unexpected error: got %s, want %s ", result, expectedResult)
	}
}
