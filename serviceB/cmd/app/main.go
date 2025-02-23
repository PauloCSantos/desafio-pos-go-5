package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"cepgraus/internal/clients"
	"cepgraus/internal/routes"
	"cepgraus/internal/tracing"
)

func main() {
	shutdown := tracing.InitTracer("service-B", "http://zipkin:9411/api/v2/spans")
	defer shutdown(context.Background())

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading file .env:", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in the environment")
	}

	cityClient := clients.NewCityClient()
	temperatureClient := clients.NewTemperatureClient(apiKey)

	router := routes.SetupRoutes(cityClient, temperatureClient)

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Erro ao iniciar o servidor", err)
	}
}
