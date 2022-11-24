package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Expires     time.Time
}

type SpotifyData struct {
	AccessToken  *AccessData
	UserToken    *TokenData
	ClientId     string
	ClientSecret string
	Dc           string
}

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Error loading .env")
	}
	// accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	sp_dc := os.Getenv("SPOTIFY_DC")

	// track_id := "5f8eCNwTlr0RJopE9vQ6mB"
	client := SpotifyData{
		ClientId:     client_id,
		ClientSecret: client_secret,
		Dc:           sp_dc,
	}
	accessInfo := client.getAccessToken()
	client.AccessToken = accessInfo
	clientInfo := client.getClientCredentialsLogin()
	clientInfo.Expires = time.Now().Add(time.Second * time.Duration(clientInfo.ExpiresIn))
	client.UserToken = clientInfo
	fmt.Println(client.UserToken.AccessToken)
	client.getTrackId()

	// client.getLyrics(track_id)
}
