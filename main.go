package main

import (
	"fiqri_muhammad/student-api/config"
	"fiqri_muhammad/student-api/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	router := routes.SetupRouter()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
