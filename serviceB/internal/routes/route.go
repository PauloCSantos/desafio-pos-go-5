package routes

import (
	"cepgraus/internal/clients"
	"cepgraus/internal/handlers"
	"net/http"
)

func SetupRoutes(cityClient *clients.CityClient, temperatureClient *clients.TemperatureClient) *http.ServeMux {
	mux := http.NewServeMux()
	handler := handlers.NewTemperatureHandler(cityClient, temperatureClient)
	mux.HandleFunc("/temperatureByCEP", handler.TemperatureByCep)

	return mux
}
