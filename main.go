package main

import (
	"log"

	"ynbauth/database"
	"github.com/joho/godotenv"
	"ynbauth/router"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDatabase()

	// Initialize router
	r := router.SetupRouter()

	// Start the server
	r.Run(":8080")
}
