package web

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/rs/cors"
	"github.com/zmb3/spotify/v2"
	iSpotify "skafteresort.se/lyrics_strips/internal/spotify"
)

// var (
// 	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
// 	ch    = make(chan *spotify.Client)
// 	state = "abc123"
// 	// These should be randomly generated for each request
// 	//  More information on generating these can be found here,
// 	// https://www.oauth.com/playground/authorization-code-with-pkce.html
// 	codeVerifier  = "w0HfYrKnG8AihqYHA9_XUPTIcqEXQvCQfOF2IitRgmlF43YWJ8dy2b49ZUwVUOR.YnvzVoTBL57BwIhM4ouSa~tdf0eE_OmiMC_ESCcVOe7maSLIk9IOdBhRstAxjCl7"
// 	codeChallenge = "ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig"
// )

func NewServer(
	logger *slog.Logger,
	allowedOrigins []string,
	jwtSecret string,

	spotifyService *iSpotify.Service,
) http.Handler {

	// panic(allowedOrigins)
	corsMw := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		Debug:          false,
		AllowedHeaders: []string{"*"},
		// AllowCredentials: true,
	})

	mux := http.NewServeMux()
	mux.Handle("/",
		corsMw.Handler(
			loggingMiddleware(logger,
				addRoutes(logger, spotifyService),
			),
		),
	)
	return mux
}

func addRoutes(
	logger *slog.Logger,
	spotifyService *iSpotify.Service,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(
		"/search",
		handleSearch(logger, spotifyService),
	)

	// mux.Handle(
	// 	"/callback",
	// 	handleCallback(logger, spotifyService),
	// )
	return mux
}

func handleSearch(
	logger *slog.Logger,
	service *iSpotify.Service,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var searchData iSpotify.SearchData
			err := json.NewDecoder(r.Body).Decode(&searchData)
			if err != nil {
				logger.Error("handleSearch", "err", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			vals := r.URL.Query()
			total := 200
			if p := vals.Get("total"); p != "" {
				total, err = strconv.Atoi(p)
				if err != nil {
					logger.Error("handleSearch", "err", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			}
			res, err := service.HandleSearch(r.Context(), searchData, total)
			if err != nil {
				logger.Error("handleSearch", "err", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var lyrics *iSpotify.LyricsLines
			var track spotify.FullTrack
			tracks := res.Tracks.Tracks
			totalTracks := res.Tracks.Total

			for {
				rInt := 1
				if len(tracks) > 1 {
					rInt = rand.Intn(len(tracks) - 1)
				}
				track = tracks[rInt]
				lyrics, err = service.GetLyrics(track)
				if err != nil {
					if err.Error() == "MISSING_LYRICS" && len(tracks) > 1 {
						localTracks := []spotify.FullTrack{}
						localTracks = tracks[:rInt]
						if rInt != len(tracks)-1 {
							localTracks = append(localTracks, tracks[rInt+1:]...)
						}
						tracks = localTracks
						continue
					}
					logger.Error("handleSearch/GetLyrics", "err", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				break
			}
			lineNumber := 0
			if len(lyrics.SyncedLyricsSlice) > 0 {
				lineNumber = rand.Intn(len(lyrics.SyncedLyricsSlice) - 1)
			}
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]any{
				"track":      track,
				"lyrics":     lyrics.SyncedLyricsSlice,
				"lineNumber": lineNumber,
				"total":      totalTracks,
			})
		},
	)
}

// func handleCallback(
// 	logger *slog.Logger,
// 	service *iSpotify.SpotifyService,
// ) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			tok, err := auth.Token(r.Context(), state, r,
// 				oauth2.SetAuthURLParam("code_verifier", codeVerifier))
// 			if err != nil {
// 				http.Error(w, "Couldn't get token", http.StatusForbidden)
// 				log.Fatal(err)
// 			}
// 			if st := r.FormValue("state"); st != state {
// 				http.NotFound(w, r)
// 				log.Fatalf("State mismatch: %s != %s\n", st, state)
// 			}
// 			// use the token to get an authenticated client
// 			client := spotify.New(auth.Client(r.Context(), tok))
// 			fmt.Fprintf(w, "Login Completed!")
// 			ch <- client
// 		},
// 	)
// }
