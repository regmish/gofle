package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func HttpGet(url string, body interface{}) (*http.Response, error) {
	client := createClient()

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	if _, exists := os.LookupEnv("AUTHORIZATION_TOKEN"); exists {
		req.Header.Add("Authorization", os.Getenv("AUTHORIZATION_TOKEN"))
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return resp, nil
}

func HttpPost(url string, body []byte) (*http.Response, error) {
	client := createClient()

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	if _, exists := os.LookupEnv("AUTHORIZATION_TOKEN"); exists {
		req.Header.Add("Authorization", os.Getenv("AUTHORIZATION_TOKEN"))
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return resp, nil
}

func Post(url string, body []byte, resp interface{}) {
	httpResp, err := HttpPost(url, body)

	if err != nil {
		panic(err.Error())
	}

	defer httpResp.Body.Close()

	json.NewDecoder(httpResp.Body).Decode(&resp)
}

func Get(url string, resp interface{}) {
	httpResp, err := HttpGet(url, nil)

	if err != nil {
		panic(err.Error())
	}

	defer httpResp.Body.Close()

	json.NewDecoder(httpResp.Body).Decode(&resp)
}

func createClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{Transport: tr}

	return httpClient
}
