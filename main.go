package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
)

type URL struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	ShortURL  string `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}

var db = make(map[string]URL)

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