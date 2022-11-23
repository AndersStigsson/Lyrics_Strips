package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (sp *SpotifyData) getLyrics(track_id string) {
	url := fmt.Sprintf("https://spclient.wg.spotify.com/color-lyrics/v2/track/%s?format=json&market=from_token", track_id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("App-platform", "WebPlayer")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", sp.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
