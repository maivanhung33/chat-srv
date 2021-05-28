package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

	port := os.Getenv("PORT")
	ServiceAccountUri = os.Getenv("SERVICE_AUTH_URI")

	http.HandleFunc("/websocket", handleConnections)
	go handleMessages()

	log.Println("[STATUS] Server starting at localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("[STATUS] Server starting at localhost:"+port+"failed", err)
	}
}
