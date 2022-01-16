package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const TINY_URL_EP = "https://api.tinyurl.com/create?api_token=LmymPWdUZtaeNw7pXcmj10GMnghUhs7r0ekXNjCrcwUbOOJRCx55yXwCSI2K"

type TinyUrlRequest struct {
	URL string `json:"url"`
}

type TinyUrlResponse struct {
	Data struct {
		URL     string        `json:"url"`
		Domain  string        `json:"domain"`
		Alias   string        `json:"alias"`
		Tags    []interface{} `json:"tags"`
		TinyURL string        `json:"tiny_url"`
	} `json:"data"`
	Code   int           `json:"code"`
	Errors []interface{} `json:"errors"`
}

func FetchTinyUrl(url string) string {
	log.Println("Inside FetchTinyUrl == ", url)
	// Create the Request Body
	tinyReq := &TinyUrlRequest{
		URL: url,
	}
	var reqBuffer bytes.Buffer
	json.NewEncoder(&reqBuffer).Encode(tinyReq)

	// Init API Client
	client := &http.Client{}
	req, err := http.NewRequest("POST", TINY_URL_EP, &reqBuffer)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	log.Println("Inside FetchTinyUrl == 1")
	// hit API
	resp, err := client.Do(req)
	log.Println("Inside FetchTinyUrl == 2")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	// parse Body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	log.Println("Inside FetchTinyUrl == 3")
	if err != nil {
		log.Println(err.Error())
	}
	var responseObject TinyUrlResponse
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Println(responseObject)
	// fetch tinyUrl
	if resp.StatusCode == http.StatusOK {
		return responseObject.Data.TinyURL
	}
	return ""
}
