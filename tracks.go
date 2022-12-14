package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Tracks struct {
	Id     string       `json:"id"`
	Name   string       `json:"name"`
	Album  AlbumData    `json:"album"`
	Artist []ArtistData `json:"artists"`
}

type TrackData struct {
	Tracks []Tracks `json:"tracks"`
}

func (sp *SpotifyData) getTrackId(values url.Values) *TrackData {
	genres := "swedish"
	if values.Has("seed_genres") {
		genres = values.Get("seed_genres")
	}
	url := fmt.Sprintf("https://api.spotify.com/v1/recommendations?seed_genres=%s&limit=10&min_popularity=33", genres)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", sp.UserToken.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data TrackData
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err.Error())
	}
	return &data
}
