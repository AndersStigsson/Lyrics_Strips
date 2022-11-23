package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Error loading .env")
	}
	access_token := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	getLyrics(access_token)
}
