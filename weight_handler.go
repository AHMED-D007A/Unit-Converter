package main

import (
	"context"
	"net/http"
	"strconv"
	"unit_converter/templates"
)

func getWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("weight", WeightUnits)
		componant.Render(context.Background(), w)
	}
}

func postWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		weight := r.FormValue("weight")
		from := r.FormValue("from")
		to := r.FormValue("to")
		result := "0"
		if weight == "" {
			weight = "0"
		}
		var wc WeightConverter
		wc.input, _ = strconv.ParseFloat(weight, 64)
		wc.from = from
		wc.to = to
		wc.initWeightConverter()
		result = strconv.FormatFloat(wc.result, 'f', -1, 64)
		componant := templates.Result(weight, from, result, to)
		componant.Render(context.Background(), w)
	}
}
