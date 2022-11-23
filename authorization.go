package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	DEFAULT_HOST   = "accounts.spotify.com"
	DEFAULT_PORT   = 443
	DEFAULT_SCHEME = "https"
)

var AccessToken string

func (sp *SpotifyData) getAccessToken() {
	url := "https://open.spotify.com"
	method := http.MethodGet

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	cookie := &http.Cookie{
		Name:   "sp_dc",
		Value:  sp.Dc,
		MaxAge: 300,
	}
	req.AddCookie(cookie)
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
	// Get the access_token value within
	// <script id="session" data-testid="session" type="application/json">
	fmt.Println(string(body))
}

func getClientCredentialsLogin(client_id string, client_secret string) {
	url := fmt.Sprintf("%s://%s/api/token", DEFAULT_SCHEME, DEFAULT_HOST)
	// url := "https://accounts.spotify.com/api/token"
	method := http.MethodPost

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	authBearer := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client_id, client_secret)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authBearer))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
