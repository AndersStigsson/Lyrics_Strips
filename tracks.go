package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (sp *SpotifyData) getTrackId() *string {
	url := "https://api.spotify.com/v1/recommendations?seed_genres=punk&limit=1&min_popularity=75"

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

	fmt.Println(string(body))
	return nil
}
