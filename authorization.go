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

func test1() {
	url := "https://accounts.spotify.com/api/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	// req.Header.Add("Authorization", "Basic ")
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

func getClientCredentialsLogin(client_id string, client_secret string) {
	requestURL := fmt.Sprintf("%s://%s/api/token", DEFAULT_SCHEME, DEFAULT_HOST)
	fmt.Println(requestURL)
	payload := strings.NewReader("grant_type=client_credentials")

	request, err := http.NewRequest(http.MethodPost, requestURL, payload)
	if err != nil {
		fmt.Println("Failed in new request")
		fmt.Println(err.Error())
	}

	authBearer := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client_id, client_secret)))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authBearer))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Print(request)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Failed when executing request")
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
