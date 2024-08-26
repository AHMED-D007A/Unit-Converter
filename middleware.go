package main

import (
	"log"
	"net/http"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
