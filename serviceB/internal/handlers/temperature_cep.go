package handlers

import (
	"cepgraus/internal/clients"
	"cepgraus/internal/services"
	"cepgraus/internal/tracing"
	"cepgraus/internal/utils"
	"encoding/json"
	"math"
	"net/http"
)

type TemperatureHandler struct {
	CityClient        *clients.CityClient
	TemperatureClient *clients.TemperatureClient
}

func NewTemperatureHandler(cityClient *clients.CityClient, temperatureClient *clients.TemperatureClient) *TemperatureHandler {
	return &TemperatureHandler{
		CityClient:        cityClient,
		TemperatureClient: temperatureClient,
	}
}

func (th *TemperatureHandler) TemperatureByCep(w http.ResponseWriter, r *http.Request) {
	//ctx, span := tracing.StartSpanFromRequest(r, "Service B: TemperatureByCep")
	ctx, span := tracing.StartSpanFromRequest(r, "Handler: TemperatureByCep")
	defer span.End()

	cep := r.URL.Query().Get("cep")

	if !services.ValidateCep(cep) {
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cityCtx, citySpan := tracing.StartSpan(ctx, "Get CityByCEP")
	cityName, err := th.CityClient.GetCityByCEP(cityCtx, cep)
	if err != nil {
		citySpan.End()
		http.Error(w, "Can not find zipcode", http.StatusNotFound)
		return
	}
	citySpan.End()

	cityNameFormatted := utils.Formatter(cityName.Localidade)

	tempCtx, tempSpan := tracing.StartSpan(ctx, "Get Temperature")
	temperature, err := th.TemperatureClient.GetTemperature(tempCtx, cityNameFormatted)
	if err != nil {
		tempSpan.End()
		http.Error(w, "Can not find temperature", http.StatusBadRequest)
		return
	}
	tempSpan.End()

	temperatureConverter, err := utils.NewTemperatureConverter(temperature, "C")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"city":   cityName.Localidade,
		"temp_C": math.Round(temperatureConverter.ToCelsius()*10) / 10,
		"temp_F": math.Round(temperatureConverter.ToFahrenheit()*10) / 10,
		"temp_K": math.Round(temperatureConverter.ToKelvin()*10) / 10,
	}

	_, writeSpan := tracing.StartSpan(ctx, "Write Response")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	writeSpan.End()
}
