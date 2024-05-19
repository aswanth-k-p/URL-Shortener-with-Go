package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/router"
	"url-shortener/storage"

	"github.com/joho/godotenv"
)

func main() {
	// MongoDB connection details
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := "urlshortener"
	collectionName := "urls"

	store := storage.NewMongoStorage(mongoURI, dbName, collectionName)
	r := router.NewRouter(store)

	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
