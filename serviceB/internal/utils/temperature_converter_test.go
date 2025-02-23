package utils

import (
	"math"
	"testing"
)

var testDatas = []struct {
	value     float64
	unit      string
	expectedF float64
	expectedC float64
	expectedK float64
}{
	{value: 16.0, unit: "C", expectedF: 60.8, expectedC: 16.0, expectedK: 289.00},
	{value: 16.0, unit: "K", expectedF: -430.87, expectedC: -257.15, expectedK: 16.0},
	{value: 16.0, unit: "F", expectedF: 16.0, expectedC: -8.8889, expectedK: 264.2611},
}

const floatTolerance = 0.001

func floatsAreEqual(a, b float64) bool {
	return math.Abs(a-b) <= floatTolerance
}

func TestInvalidUnitError(t *testing.T) {
	value := 16.0
	unit := "X"

	_, err := NewTemperatureConverter(value, unit)

	if err == nil {
		t.Errorf("Expected error for invalid unit %q, got nil", unit)
	}
}

func TestConversionToFahrenheit(t *testing.T) {
	for _, testData := range testDatas {
		converter, err := NewTemperatureConverter(testData.value, testData.unit)
		if err != nil {
			t.Errorf("Unexpected error for valid unit %q: %v", testData.unit, err)
			continue
		}

		resultF := converter.ToFahrenheit()
		if !floatsAreEqual(resultF, testData.expectedF) {
			t.Errorf("ToFahrenheit failed for unit %q: got %f, want %f", testData.unit, resultF, testData.expectedF)
		}
	}
}

func TestConversionToCelsius(t *testing.T) {
	for _, testData := range testDatas {
		converter, err := NewTemperatureConverter(testData.value, testData.unit)
		if err != nil {
			t.Errorf("Unexpected error for valid unit %q: %v", testData.unit, err)
			continue
		}

		resultC := converter.ToCelsius()
		if !floatsAreEqual(resultC, testData.expectedC) {
			t.Errorf("ToCelsius failed for unit %q: got %f, want %f", testData.unit, resultC, testData.expectedC)
		}
	}
}

func TestConversionToKelvin(t *testing.T) {
	for _, testData := range testDatas {
		converter, err := NewTemperatureConverter(testData.value, testData.unit)
		if err != nil {
			t.Errorf("Unexpected error for valid unit %q: %v", testData.unit, err)
			continue
		}

		resultK := converter.ToKelvin()
		if !floatsAreEqual(resultK, testData.expectedK) {
			t.Errorf("ToKelvin failed for unit %q: got %f, want %f", testData.unit, resultK, testData.expectedK)
		}
	}
}
