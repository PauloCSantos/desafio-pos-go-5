package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCityByCEP_Success(t *testing.T) {
	response := `{
		"cep": "13720-616",
		"logradouro": "Rua Lagoa da Pampulha",
		"bairro": "Residencial dos Lagos",
		"localidade": "São José do Rio Pardo",
		"uf": "SP",
		"estado": "São Paulo"
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	client := NewCityClient()
	client.BaseURL = ts.URL + "/"

	// Criando um contexto para o teste
	ctx := context.Background()
	city, err := client.GetCityByCEP(ctx, "13720-616")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedCity := &City{
		Cep:        "13720-616",
		Logradouro: "Rua Lagoa da Pampulha",
		Bairro:     "Residencial dos Lagos",
		Localidade: "São José do Rio Pardo",
		UF:         "SP",
		Estado:     "São Paulo",
	}
	if *city != *expectedCity {
		t.Fatalf("expected %+v, got %+v", expectedCity, city)
	}
}

func TestGetCityByCEP_NotFound(t *testing.T) {
	response := `{"erro": true}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	client := NewCityClient()
	client.BaseURL = ts.URL + "/"

	// Criando um contexto para o teste
	ctx := context.Background()
	city, err := client.GetCityByCEP(ctx, "00000-000")
	if err == nil {
		t.Fatalf("expected an error, got none")
	}

	if city != nil {
		t.Fatalf("expected nil city, got %+v", city)
	}
}
