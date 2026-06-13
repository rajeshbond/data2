package main

import (
	"log"
	"net/http"
	"os"

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

	// Render provides PORT dynamically
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local testing
	}

	log.Println("Server Started On :" + port)

	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/rajbond/data2/internal/handlers"
// )

// func main() {

// 	http.HandleFunc(
// 		"/api/v1/cycle",
// 		handlers.CreateCycle,
// 	)

// 	http.HandleFunc(
// 		"/api/v1/cycle/latest",
// 		handlers.GetLatestCycle,
// 	)

// 	log.Println(
// 		"Server Started On :8080",
// 	)

// 	log.Fatal(
// 		http.ListenAndServe(
// 			":8080",
// 			nil,
// 		),
// 	)
// }
