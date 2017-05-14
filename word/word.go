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

var translateUrl = "https://translate.yandex.net/api/v1.5/tr.json/translate"
var getSupportedLanguagesUrl = "https://translate.yandex.net/api/v1.5/tr.json/getLangs"
var api_key = os.Getenv("TRANSLATION_API_KEY")
var supportedLanguages = getSupportedLanguages()

func main() {
    http.HandleFunc("/word/", translationHandler)
    http.ListenAndServe(":8080", nil)

}

func translationHandler(writer http.ResponseWriter, request *http.Request) {
    language := request.URL.Path[6:]
    if len(language) == 0 {
        fmt.Fprintf(writer, "Please select a language to translate to. The language codes available are: \n")
        for key, value := range supportedLanguages {
            fmt.Fprintf(writer, "%s - %s\n", key, value)
        }
        return
    }
    word := generator.GenerateRandomWord()
    req := buildTranslateRequest(word, language)
    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    apiResponse := handleResponse(resp, writer, language)

    for _, translatedWord := range apiResponse.Text {
        fmt.Fprintf(writer, "Um tradução de \"%s\" é: %s \n", word, translatedWord)
    }
}

func buildTranslateRequest(word string, language string) *http.Request {
    form := url.Values{}
    form.Add("key", api_key)
    form.Add("text", word)
    form.Add("lang", "en-" + language)

    req, err := http.NewRequest("POST", translateUrl, strings.NewReader(form.Encode()))
    if err != nil {
        fmt.Println("Something went wrong building the HTTP request")
        panic(err)
    }
    req.PostForm = form
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    return req
    
}

func handleResponse(resp *http.Response, writer http.ResponseWriter, language string) *Translation {
    defer resp.Body.Close()

    translation := new(Translation)
    err := json.NewDecoder(resp.Body).Decode(translation)

    if(err != nil){
        fmt.Println("Error unmarshalling:", err)
    }
    if translation.Code != 200 {
        fmt.Println("Error code received from the translation API: ", 
            strconv.Itoa(translation.Code) + " - " + errorResponseCodes[translation.Code])
        if translation.Code == 501 {
            fmt.Fprintf(writer, "%s is not an accepted language code to translate. The language codes available are: \n", language)
            for key, value := range supportedLanguages {
                fmt.Fprintf(writer, "%s - %s\n", key, value)
            }
        } else {
            fmt.Fprintf(writer, "There was an error processing your request. Please try again.")
        }
    }
    return translation
}

func getSupportedLanguages() map [string]string{
    form := url.Values{}
    form.Add("key", api_key)
    form.Add("ui", "en")

    req, err := http.NewRequest("POST", getSupportedLanguagesUrl, strings.NewReader(form.Encode()))
    if err != nil {
        fmt.Println("Something went wrong building the HTTP request")
        panic(err)
    }
    req.PostForm = form
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := http.Client{}
    resp, err := client.Do(req)

    languages := new(Languages)
    err = json.NewDecoder(resp.Body).Decode(languages)

    if(err != nil){
        fmt.Println("Error unmarshalling:", err)
    }

    resp.Body.Close()
    return languages.Langs
} 

type Translation struct {
    Code int `json:"code"`
    Lang string `json:"lang"`
    Text []string `json:"text"`
}

type Languages struct {
    Dirs []string `json:"dirs"`
    Langs map[string] string `json:"langs"`
}

var errorResponseCodes = map[int]string {
    401 : "Invalid API key",
    402 : "Blocked API key",
    404 : "Exceeded the daily limit on the amount of translated text",
    413 : "Exceeded the maximum text size",
    422 : "The text cannot be translated",
    501 : "The specified translation direction is not supported" }
