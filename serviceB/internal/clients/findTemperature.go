package clients

import (
	"cepgraus/internal/tracing"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Temperature struct {
	Temp float64 `json:"temp_c"`
}

type WeatherResponse struct {
	Temperature Temperature `json:"current"`
}

type TemperatureClient struct {
	BaseURL string
	Client  *http.Client
	apiKey  string
}

func NewTemperatureClient(apiKey string) *TemperatureClient {
	return &TemperatureClient{
		BaseURL: "http://api.weatherapi.com/v1",
		Client:  &http.Client{Timeout: 10 * time.Second},
		apiKey:  apiKey,
	}
}

func (t *TemperatureClient) GetTemperature(ctx context.Context, cityName string) (float64, error) {
	ctx, span := tracing.StartSpan(ctx, "client: GetTemperature")
	defer span.End()

	cityNameEscaped := url.QueryEscape(cityName)
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", t.BaseURL, t.apiKey, cityNameEscaped)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}

	tracing.InjectTraceIntoRequest(ctx, req)

	req.Header.Set("Content-Type", "application/json")

	resp, err := t.Client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&apiError); err == nil {
			return 0, fmt.Errorf("API error: %s (code: %d)", apiError.Error.Message, apiError.Error.Code)
		}

		return 0, fmt.Errorf("failed to get city temperature, status: %s", resp.Status)
	}

	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return 0, fmt.Errorf("error decoding response: %w", err)
	}

	return weatherResponse.Temperature.Temp, nil
}
