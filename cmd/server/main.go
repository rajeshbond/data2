package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rajbond/data2/internal/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	// Router
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/api/v1/cycle", handlers.CreateCycle)
	mux.HandleFunc("/api/v1/cycle/latest", handlers.GetLatestCycle)

	// Render dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server Started On :" + port)

	// Wrap router with CORS middleware
	err := http.ListenAndServe(":"+port, enableCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}
