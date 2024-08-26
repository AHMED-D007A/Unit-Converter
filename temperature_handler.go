package main

import (
	"context"
	"net/http"
	"strconv"
	"unit_converter/templates"
)

func getTemperature() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("temperature", TemperatureUnits)
		componant.Render(context.Background(), w)
	}
}

func postTemperature() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		temperature := r.FormValue("temperature")
		from := r.FormValue("from")
		to := r.FormValue("to")
		result := "0"
		if temperature == "" {
			temperature = "0"
		}
		var tc TemperatureConverter
		tc.input, _ = strconv.ParseFloat(temperature, 64)
		tc.from = from
		tc.to = to
		tc.initTemperatureConverter()
		result = strconv.FormatFloat(tc.result, 'f', -1, 64)
		componant := templates.Result(temperature, from, result, to)
		componant.Render(context.Background(), w)
	}
}
