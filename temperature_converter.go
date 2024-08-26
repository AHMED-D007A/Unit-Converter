package main

var TemperatureUnits = []string{"celsius", "fahrenheit", "kelvin"}

type TemperatureConverter struct {
	input  float64
	from   string
	to     string
	result float64

	clsius     float64
	fahrenheit float64
	kelvin     float64
}

func (tc *TemperatureConverter) initTemperatureConverter() {
	switch tc.from {
	case "celsius":
		tc.clsius = tc.input
	case "fahrenheit":
		tc.fahrenheit = tc.input
		tc.from_Fahrenheit()
	case "kelvin":
		tc.kelvin = tc.input
		tc.from_Kelvin()
	}

	switch tc.to {
	case "celsius":
		tc.result = tc.clsius
	case "fahrenheit":
		tc.to_Fahrenheit()
		tc.result = tc.fahrenheit
	case "kelvin":
		tc.to_Kelvin()
		tc.result = tc.kelvin
	}
}

func (tc *TemperatureConverter) to_Fahrenheit() { tc.fahrenheit = (tc.clsius * 9 / 5) + 32 }
func (tc *TemperatureConverter) to_Kelvin()     { tc.kelvin = tc.clsius + 273.15 }

func (tc *TemperatureConverter) from_Fahrenheit() { tc.clsius = (tc.fahrenheit - 32) * 5 / 9 }
func (tc *TemperatureConverter) from_Kelvin()     { tc.clsius = tc.kelvin - 273.15 }
