package clients

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"cepgraus/internal/tracing"
)

type ServiceBClient struct {
	BaseURL string
	Client  *http.Client
}

func NewServiceBClient() *ServiceBClient {
	return &ServiceBClient{
		BaseURL: "http://serviceB:8080/temperatureByCEP",
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *ServiceBClient) GetTemperatureByCEP(ctx context.Context, cep string) ([]byte, error) {
	ctx, span := tracing.StartSpan(ctx, "client: Call Service B")
	defer span.End()

	url := fmt.Sprintf("%s?cep=%s", s.BaseURL, cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	tracing.InjectTraceIntoRequest(ctx, req)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get temperature, status: %s, body: %s", resp.Status, string(errorBody))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
