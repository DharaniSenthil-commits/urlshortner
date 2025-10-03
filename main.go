package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var urlStore = make(map[string]string)
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateShortId() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func shortnerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}
	var body struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	short := generateShortId()
	urlStore[short] = body.URL

	resp := map[string]string{
		"short": short,
		"url":   body.URL,
	}

	json.NewEncoder(w).Encode(resp)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Path[1:]
	if url, ok := urlStore[short]; ok {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/shorten", shortnerHandler)
	http.HandleFunc("/", redirectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
