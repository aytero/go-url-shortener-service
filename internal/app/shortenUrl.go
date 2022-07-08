package app

import (
	"math/rand"
)

const (
	urlLength      = 10
	urlBaseSymbols = "_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = urlBaseSymbols[rand.Intn(len(urlBaseSymbols))]
	}
	return string(b)
}

func (ap *App) makeHash() string {
	//hdm5Url := md5.Sum([]byte(fullUrl))
	hash := randStringBytes(urlLength)
	return hash
}

func (ap *App) ShortenUrl(fullUrl string) (string, error) {
	ap.db.Lock()
	foundEntry, err := ap.db.GetShortUrl(fullUrl)
	ap.db.Unlock()
	if err == nil {
		ap.cfg.InfoLog.Println("Url already exists")
		return foundEntry, nil
	}
	shortUrl := ap.makeHash()
	unique := false
	ap.db.Lock()
	for !unique {
		foundEntry, err = ap.db.GetUrl(shortUrl)
		if err != nil {
			unique = true
		}
		shortUrl = ap.makeHash()
	}
	ap.db.Unlock()

	ap.db.Lock()
	_, err = ap.db.PostUrl(shortUrl, fullUrl)
	ap.db.Unlock()
	if err != nil {
		ap.cfg.ErrorLog.Println(err)
		return "", err
	}
	return ap.cfg.UrlHost + shortUrl, nil
}

func (ap *App) GetFullUrl(shortUrl string) (string, error) {
	ap.db.Lock()
	fullUrl, err := ap.db.GetUrl(shortUrl)
	ap.db.Unlock()
	if err != nil {
		ap.cfg.ErrorLog.Println(err)
		return "", err
	}
	return fullUrl, nil
}
