package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cepgraus/internal/clients"
	"cepgraus/internal/routes"
	"cepgraus/internal/tracing"
)

func main() {
	shutdown := tracing.InitTracer("service-A", "http://zipkin:9411/api/v2/spans")
	defer shutdown(context.Background())

	serviceBClient := clients.NewServiceBClient()
	router := routes.SetupRoutes(serviceBClient)

	fmt.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("Erro ao iniciar o servidor", err)
	}
}
