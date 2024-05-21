package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

//Memory DB
type URL struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	ShortURL  string `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}

var db = make(map[string]URL)


//Shortening functions

func makeShortURL(url string) string {
	hash := md5.New()
	hash.Write([]byte(url))

	returnHash:= hex.EncodeToString(hash.Sum(nil))

	return returnHash[:8]
}

func createURL(url string) string {
	short:= makeShortURL(url)

	id:= short
	db[id] = URL{
		ID: id,
		URL: url,
		ShortURL: short,
		CreatedAt: time.Now(),
	}

	return short
}

func getURL(id string) (URL, error) {
	url, ok := db[id]

	if !ok {
		return URL{}, errors.New("Not found")
	}

	return url, nil
}



// HTTP Handlers

func page(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err:= json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	URLShorter:= createURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: URLShorter}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	id:= r.URL.Path[len("/redirect"):]
	if id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.ShortURL, http.StatusFound)
}