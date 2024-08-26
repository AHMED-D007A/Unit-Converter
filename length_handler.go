package main

import (
	"context"
	"net/http"
	"strconv"
	"unit_converter/templates"
)

func getLength() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("length", LengthUnits)
		componant.Render(context.Background(), w)
	}
}

func postLength() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		length := r.FormValue("length")
		from := r.FormValue("from")
		to := r.FormValue("to")
		result := "0"
		if length == "" {
			length = "0"
		}
		var lc LengthConverter
		lc.input, _ = strconv.ParseFloat(length, 64)
		lc.from = from
		lc.to = to
		lc.initLengthConverter()
		result = strconv.FormatFloat(lc.result, 'f', -1, 64)
		componant := templates.Result(length, from, result, to)
		componant.Render(context.Background(), w)
	}
}
