package main

import (
	"fmt"
	"random-word-translator/generator"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"os"
	"strconv"
)

var translatorAPI = "https://translate.yandex.net/api/v1.5/tr.json/translate"
var api_key = os.Getenv("TRANSLATION_API_KEY")
// var api_key = "trnsl.1.1.20170513T204309Z.5861366a1aa3a27b.f3ebfce9182383782687fab1a4ae61432f3fa86b"
var language = "pt"

func main() {
	word := generator.GenerateRandomWord()
	request := buildRequest(word, language)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
    apiResponse := handleResponse(resp)

    fmt.Println("A tradução de \"" + word + "\" é: " + apiResponse.Text[0] + "\n")
}

func buildRequest(word string, language string) *http.Request {
	form := url.Values{}
    form.Add("key", api_key)
    form.Add("text", word)
    form.Add("lang", "en-" + language)

    req, err := http.NewRequest("POST", translatorAPI, strings.NewReader(form.Encode()))
    if err != nil {
    	fmt.Println("Something went wrong building the HTTP request")
        panic(err)
    }
    req.PostForm = form
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    return req
    
}

func handleResponse(resp *http.Response) *Translation {
	defer resp.Body.Close()

    translation := new(Translation)
    err := json.NewDecoder(resp.Body).Decode(translation)

    if(err != nil){
        fmt.Println("Error unmarshalling:", err)
    }
    if translation.Code != 200 {
    	fmt.Println("Error code received from the translation API: ", 
    		strconv.Itoa(translation.Code) + " - " + errorResponseCodes[translation.Code])
    	os.Exit(1)
    }
    return translation
}

type Translation struct {
    Code int `json:"code"`
    Lang string `json:"lang"`
    Text []string `json:"text"`
}

var errorResponseCodes = map[int]string {
	401 : "Invalid API key",
	402	: "Blocked API key",
	404	: "Exceeded the daily limit on the amount of translated text",
	413	: "Exceeded the maximum text size",
	422	: "The text cannot be translated",
	501	: "The specified translation direction is not supported"}
