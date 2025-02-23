package clients

import (
	"cepgraus/internal/tracing"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type City struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Erro        bool   `json:"erro"`
}

type CityClient struct {
	BaseURL string
	Client  *http.Client
}

func NewCityClient() *CityClient {
	return &CityClient{
		BaseURL: "http://viacep.com.br/ws/",
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *CityClient) GetCityByCEP(ctx context.Context, cep string) (*City, error) {
	ctx, span := tracing.StartSpan(ctx, "client: GetCityByCEP")
	defer span.End()

	url := fmt.Sprintf("%s%s/json/", c.BaseURL, cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	tracing.InjectTraceIntoRequest(ctx, req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get city, status: %s", resp.Status)
	}

	var city City
	err = json.NewDecoder(resp.Body).Decode(&city)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if city.Erro {
		return nil, fmt.Errorf("CEP not found")
	}

	return &city, nil
}
