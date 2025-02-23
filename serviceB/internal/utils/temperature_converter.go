package utils

import (
	"fmt"
)

type temperatureConverter struct {
	value float64
	unit  string
}

type TemperatureConverter interface {
	ToFahrenheit() float64
	ToCelsius() float64
	ToKelvin() float64
}

func NewTemperatureConverter(value float64, unit string) (TemperatureConverter, error) {
	validUnits := map[string]bool{"C": true, "F": true, "K": true}
	if !validUnits[unit] {
		return nil, fmt.Errorf("invalid temperature unit: %s", unit)
	}
	return &temperatureConverter{value: value, unit: unit}, nil
}

func (t *temperatureConverter) ToFahrenheit() float64 {
	switch t.unit {
	case "C":
		return t.value*1.8 + 32
	case "K":
		return (t.value-273.15)*9/5 + 32
	case "F":
		return t.value
	default:
		panic("Unsupported temperature unit")
	}
}

func (t *temperatureConverter) ToCelsius() float64 {
	switch t.unit {
	case "F":
		return (t.value - 32) * 5 / 9
	case "K":
		return t.value - 273.15
	case "C":
		return t.value
	default:
		panic("Unsupported temperature unit")
	}
}

func (t *temperatureConverter) ToKelvin() float64 {
	switch t.unit {
	case "C":
		return t.value + 273
	case "F":
		return (t.value-32)*5/9 + 273.15
	case "K":
		return t.value
	default:
		panic("Unsupported temperature unit")
	}
}
