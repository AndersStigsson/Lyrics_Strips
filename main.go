package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

var client SpotifyData

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func validateRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		}
		if client.AccessToken.Expires.Before(time.Now()) {
			// get new accessToken
		}
		if client.UserToken.Expires.Before(time.Now()) {
			// get new UserToken
		}
		next(w, r)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/banan", validateRequest(addNewGuest))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-api-token"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10010", handlers.CORS(headersOk, methodsOk)(myRouter)))
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
	client = SpotifyData{
		ClientId:     client_id,
		ClientSecret: client_secret,
		Dc:           sp_dc,
	}
	accessInfo := client.getAccessToken()
	client.AccessToken = accessInfo
	clientInfo := client.getClientCredentialsLogin()
	clientInfo.Expires = time.Now().Add(time.Second * time.Duration(clientInfo.ExpiresIn))
	client.UserToken = clientInfo
	trackData := client.getTrackId()
	fmt.Println(client.UserToken.AccessToken)
	handleRequests()

	numberOfSongs := len(trackData.Tracks)
	songNumber := rand.Intn(numberOfSongs)
	lyrics := client.getLyrics(trackData.Tracks[songNumber].Id)
	lineNumber := rand.Intn(len(lyrics.Lyrics.Lines))
	lineOfText := lyrics.Lyrics.Lines[lineNumber]
	for {
		if len(strings.Split(lineOfText.Words, " ")) > 3 {
			break
		}
		lineNumber = rand.Intn(len(lyrics.Lyrics.Lines))
		lineOfText = lyrics.Lyrics.Lines[lineNumber]
	}

	fmt.Println(lineOfText)
	fmt.Scanln()
	fmt.Printf("Namn: %s, Artist: %s", trackData.Tracks[songNumber].Name, trackData.Tracks[songNumber].Artist[0].Name)
}
