package iSpotify

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type SearchData struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type Service struct {
	logger *slog.Logger
	client *spotify.Client
}

type SyncedLyrics struct {
	Time  string `json:"time"`
	Words string `json:"words"`
}

type LyricsLines struct {
	PlainLyrics       string `json:"plainLyrics"`
	SyncedLyrics      string `json:"syncedLyrics"`
	SyncedLyricsSlice []SyncedLyrics
}

// type Lyrics struct {
// 	Lines []LyricsLines `json:"lines"`
// }
//
// type LyricsData struct {
// 	Lyrics Lyrics `json:"lyrics"`
// }

func NewSpotifyService(logger *slog.Logger, client *spotify.Client) *Service {
	ts := Service{
		logger: logger,
		client: client,
	}
	return &ts
}

func (s *Service) HandleSearch(ctx context.Context, sd SearchData, page int) (*spotify.SearchResult, error) {
	query := strings.Join(sd.Values, ",")
	if query == "" {
		query = "swedish"
	}
	res, err := s.client.Search(
		ctx,
		query,
		spotify.SearchTypeTrack,
		spotify.Limit(10),
		spotify.Offset(10*(page-1)),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetLyrics(track spotify.FullTrack) (*LyricsLines, error) {
	url := fmt.Sprintf("https://lrclib.net/api/get?track_name=%s&album_name=%s&artist_name=%s&duration=%d", url.QueryEscape(track.Name), url.QueryEscape(track.Album.Name), url.QueryEscape(track.Artists[0].Name), track.Duration/1000)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Lyrics-strips v0.2.0 (https://github.com/SrSelse/Lyrics_strips)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("MISSING_LYRICS")
	}
	var lyrics LyricsLines
	err = json.NewDecoder(res.Body).Decode(&lyrics)
	if err != nil {
		return nil, err
	}
	sll := strings.Split(lyrics.SyncedLyrics, "\n")
	for _, str := range sll {
		l := SyncedLyrics{}
		re := regexp.MustCompile(`\[(.*)\](.*)`)
		matches := re.FindAllStringSubmatch(str, 1)
		if len(matches[0]) > 1 {
			l.Time = matches[0][1]
			l.Words = matches[0][2]
		}

		lyrics.SyncedLyricsSlice = append(lyrics.SyncedLyricsSlice, l)
	}

	return &lyrics, nil
}
