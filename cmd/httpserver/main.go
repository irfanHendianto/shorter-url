package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"shorter-url/internal/db"                         // Import the db package for MongoDB initialization
	deliveryHttp "shorter-url/internal/delivery/http" // Import the delivery package for HTTP handlers
	"shorter-url/internal/repository"                 // Import the repository package
	"shorter-url/internal/usecase"                    // Import the use case package

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT is not set
		fmt.Println("PORT environment variable not set, defaulting to 8080")
	}

	shortURLOptions := os.Getenv("SHORT_URL_OPTIONS")

	fmt.Printf("SHORT_URL_OPTIONS: %s\n", shortURLOptions)
	client, err := db.InitMongoDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Initialize the repository with MongoDB
	urlRepo := repository.NewMongoURLRepository(client.Database("shorter-url").Collection("urls"))

	// Initialize the use case by passing in the repository
	urlShortenerUseCase := usecase.NewShortenURL(urlRepo)

	// Initialize the HTTP handler by passing in the use case
	urlHandler := deliveryHttp.NewURLHandler(urlShortenerUseCase)

	// Define the route for shortening URLs
	http.HandleFunc("/shorten", urlHandler.ShortenURLHandler)
	http.HandleFunc("/", urlHandler.RedirectURLHandler)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}