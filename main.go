package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unit_converter/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Index()
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("GET /length", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("length", LengthUnits)
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("POST /length", func(w http.ResponseWriter, r *http.Request) {
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
	})

	mux.HandleFunc("GET /weight", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("weight", WeightUnits)
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("POST /weight", func(w http.ResponseWriter, r *http.Request) {
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
	})

	mux.HandleFunc("GET /temperature", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Convert("temperature", TemperatureUnits)
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("POST /temperature", func(w http.ResponseWriter, r *http.Request) {
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
	})

	server := http.Server{
		Addr:    ":4000",
		Handler: logging(mux),
	}

	fmt.Printf("Server is up and running on port: %v\n", server.Addr[1:])
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
