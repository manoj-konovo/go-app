package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofor-little/env"
	"github.com/integrationninjas/go-app/handlers"
)

func main() {
	loadDotEnv()

	port := getEnv("PORT", "8080")

	// Initialize logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register API handlers
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/items", handlers.ItemsHandler)
	http.HandleFunc("/user", handlers.GetUser)

	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadDotEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := env.Load(".env"); err != nil {
			log.Printf("Unable to load .env file: %v", err)
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
