package main

import (
	"log"
	"net/http"

	"github.com/rajbond/data2/internal/handlers"
)

func main() {

	http.HandleFunc(
		"/api/v1/cycle",
		handlers.CreateCycle,
	)

	http.HandleFunc(
		"/api/v1/cycle/latest",
		handlers.GetLatestCycle,
	)

	log.Println(
		"Server Started On :8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			nil,
		),
	)
}
