package main


type URL struct {
	ID string `json:"id"`
	URL string `json:"url"`
	ShortURL string `json:"short_url"`
	CreatedAt string `json:"created_at"`
}

var db = make(map[string]URL)

