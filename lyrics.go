package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LyricsLines struct {
	Words string `json:"words"`
	Time  string `json:"startTimeMs"`
}

type Lyrics struct {
	Lines []LyricsLines `json:"lines"`
}

type LyricsData struct {
	Lyrics Lyrics `json:"lyrics"`
}

func (sp *SpotifyData) getLyrics(track_id string) *LyricsData {
	url := fmt.Sprintf("https://spclient.wg.spotify.com/color-lyrics/v2/track/%s?format=json&market=from_token", track_id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("App-platform", "WebPlayer")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", sp.AccessToken.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var lyrics LyricsData
	err = json.Unmarshal(body, &lyrics)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &lyrics
}
