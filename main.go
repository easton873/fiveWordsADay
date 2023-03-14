package main

import (
	"encoding/json"
	"fmt"
	myLog "github.com/easton873/five-words-a-day/log"
	"log"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LearnResponse struct {
	Words []Word `json:"words"`
}

type Word struct {
	Language   string   `json:"language"`
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	OtherData  []string `json:"otherData"`
}

func newWord(language string, word string, definition string) Word {
	return Word{Language: language, Word: word, Definition: definition}
}

// expects localhost:8080/hello?id=jimbo
func hello(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		// TODO didn't find anything
	}
	_, err := fmt.Fprintf(w, "Hello there, %s\n", id)
	if err != nil {
		myLog.LogError("Failed to write", err)
	}
}

func register(w http.ResponseWriter, req *http.Request) {
	var data RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		myLog.LogError("Failed to decode json", err)
	}
	registerMsg := fmt.Sprintf("Registered %s\n", data.Username)
	log.Printf(registerMsg)
	fmt.Fprintf(w, registerMsg)
}

func getFiveWordsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fiveWords := LearnResponse{Words: []Word{
		newWord("zh-cmn-Hant", "nihao", "hello"),
		newWord("zh-cmn-Hant", "xiedu", "blasphemy"),
		newWord("zh-cmn-Hant", "haodi", "good ground"),
		newWord("zh-cmn-Hant", "paobu", "run"),
		newWord("zh-cmn-Hant", "dayang", "close"),
	}}
	err := json.NewEncoder(w).Encode(fiveWords)
	if err != nil {
		myLog.LogError("Failed to encode json and write it to response", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/register", register)
	http.HandleFunc("/learn", getFiveWordsHandler)

	openDB()
	log.Println("Up and at 'em")
	startServer()
}

func startServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		myLog.LogError("Failed to ListenAndServe", err)
	}
}
