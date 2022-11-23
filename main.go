package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type SpotifyData struct {
	AccessToken  string
	ClientId     string
	ClientSecret string
	Dc           string
}

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Error loading .env")
	}
	accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	sp_dc := os.Getenv("SPOTIFY_DC")

	// track_id := "5f8eCNwTlr0RJopE9vQ6mB"
	client := SpotifyData{
		AccessToken:  accessToken,
		ClientId:     client_id,
		ClientSecret: client_secret,
		Dc:           sp_dc,
	}
	client.getAccessToken()
	// client.getLyrics(track_id)
}
