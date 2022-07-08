package model

type DbEntry struct {
	Id       int    `json:"id"`
	ShortUrl string `json:"short_url"`
	FullUrl  string `json:"full_url"`
}
