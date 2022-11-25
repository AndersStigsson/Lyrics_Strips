package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	DEFAULT_HOST   = "accounts.spotify.com"
	DEFAULT_PORT   = 443
	DEFAULT_SCHEME = "https"
)

var AccessToken string

type AccessData struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"accessTokenExpirationTimestampMs"`
	Expires     time.Time
}

func (sp *SpotifyData) getAccessToken() *AccessData {
	url := "https://open.spotify.com"
	method := http.MethodGet

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
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
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	r := regexp.MustCompile(`<script id="session" data-testid="session" type="application/json">(\S+)</script>`)
	searchBody := string(body)
	match := r.FindStringSubmatch(searchBody)
	var data AccessData
	err = json.Unmarshal([]byte(match[1]), &data)
	if err != nil {
		panic(err.Error())
	}
	data.Expires = time.Now().Add(time.Second * time.Duration(data.ExpiresIn))
	return &data
}

func (sp *SpotifyData) getClientCredentialsLogin() *TokenData {
	url := fmt.Sprintf("%s://%s/api/token", DEFAULT_SCHEME, DEFAULT_HOST)
	// url := "https://accounts.spotify.com/api/token"
	method := http.MethodPost

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	authString := fmt.Sprintf("%s:%s", sp.ClientId, sp.ClientSecret)
	authBearer := base64.StdEncoding.EncodeToString([]byte(authString))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authBearer))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var tempData TokenData
	err = json.Unmarshal(body, &tempData)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &tempData
}
