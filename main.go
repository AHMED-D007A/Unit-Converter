package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"unit_converter/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Index()
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("GET /length", getLength())

	mux.HandleFunc("POST /length", postLength())

	mux.HandleFunc("GET /weight", getWeight())

	mux.HandleFunc("POST /weight", postWeight())

	mux.HandleFunc("GET /temperature", getTemperature())

	mux.HandleFunc("POST /temperature", postTemperature())

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
