package handlers

import (
	"cepgraus/internal/clients"
	"cepgraus/internal/services"
	"cepgraus/internal/tracing"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetTemperatureHandler struct {
	ServiceBClient *clients.ServiceBClient
}

type requestBody struct {
	Cep string `json:"cep"`
}

func NewGetTemperatureHandler(serviceBClient *clients.ServiceBClient) *GetTemperatureHandler {
	return &GetTemperatureHandler{
		ServiceBClient: serviceBClient,
	}
}

func (th *GetTemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracing.StartSpan(r.Context(), "Handler: GetTemperature")
	defer span.End()

	var reqBody requestBody
	_, decodeSpan := tracing.StartSpan(ctx, "Decode JSON Request")
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		decodeSpan.End()
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	decodeSpan.End()

	_, validateSpan := tracing.StartSpan(ctx, "Validate CEP")
	if !services.ValidateInput(reqBody.Cep) {
		validateSpan.End()
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	validateSpan.End()

	serviceBCtx, serviceBSpan := tracing.StartSpan(ctx, "Call Service B")
	temperatures, err := th.ServiceBClient.GetTemperatureByCEP(serviceBCtx, reqBody.Cep)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Can not find temperature", http.StatusBadRequest)
		return
	}
	serviceBSpan.End()

	_, responseSpan := tracing.StartSpan(ctx, "Write Response")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(temperatures)
	responseSpan.End()
}
