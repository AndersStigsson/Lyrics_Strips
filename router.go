package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

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
			client.AccessToken = client.getAccessToken()
		}
		if client.UserToken.Expires.Before(time.Now()) {
			client.UserToken = client.getClientCredentialsLogin()
		}
		next(w, r)
	}
}

func getNextTrack(w http.ResponseWriter, r *http.Request) {
	trackData := client.getTrackId()

	numberOfSongs := len(trackData.Tracks)
	songNumber := rand.Intn(numberOfSongs)
	lyrics := client.getLyrics(trackData.Tracks[songNumber].Id)
	lineNumber := rand.Intn(len(lyrics.Lyrics.Lines))
	lineOfText := lyrics.Lyrics.Lines[lineNumber]

	for {
		if len(strings.Split(lineOfText.Words, " ")) > 2 {
			break
		}
		lineNumber = rand.Intn(len(lyrics.Lyrics.Lines))
		lineOfText = lyrics.Lyrics.Lines[lineNumber]
	}
	type ReturnData struct {
		LineNumber int           `json:"lineNumber"`
		Lines      []LyricsLines `json:"lines"`
		Track      Tracks        `json:"track"`
	}

	var returnData ReturnData
	returnData.Lines = lyrics.Lyrics.Lines
	returnData.Track = trackData.Tracks[songNumber]
	returnData.LineNumber = lineNumber

	json.NewEncoder(w).Encode(returnData)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/next", validateRequest(getNextTrack))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-api-token"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10011", handlers.CORS(headersOk, methodsOk)(myRouter)))
}
