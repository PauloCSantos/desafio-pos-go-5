package services

import "testing"

func TestValidateCep(t *testing.T) {
	validCeps := []string{"12345678", "87654321"}
	invalidCeps := []string{
		"1234567",   // Menos de 8 caracteres
		"123456789", // Mais de 8 caracteres
		"12345a78",  // Contém letras
		"12345 678", // Contém espaço
		"",          // String vazia
	}

	for _, cep := range validCeps {
		if !ValidateCep(cep) {
			t.Errorf("Expected CEP %q to be valid, but it was not", cep)
		}
	}

	for _, cep := range invalidCeps {
		if ValidateCep(cep) {
			t.Errorf("Expected CEP %q to be invalid, but it was valid", cep)
		}
	}
}
