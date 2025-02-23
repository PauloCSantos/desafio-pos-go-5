package routes

import (
	"cepgraus/internal/clients"
	"cepgraus/internal/handlers"
	"net/http"
)

func SetupRoutes(serviceBClient *clients.ServiceBClient) *http.ServeMux {
	mux := http.NewServeMux()
	handler := handlers.NewGetTemperatureHandler(serviceBClient)
	mux.HandleFunc("/getTemperature", handler.GetTemperature)

	return mux
}
