package app

import (
	"math/rand"
	"net/url"
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
	_, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		ap.cfg.ErrorLog.Println(err)
		return "", err
	}
	//foundEntry := ap.db.SearchFullUrl(fullUrl)
	//if foundEntry != nil {
	//	ap.cfg.InfoLog.Println("Url already exists")
	//	return foundEntry.ShortUrl, nil
	//}
	shortUrl := ap.makeHash()
	//foundEntry = ap.db.SearchShortUrl(shortUrl)
	//// do in a loop
	//if foundEntry != nil {
	//	ap.cfg.InfoLog.Println("Url hash collision")
	//	shortUrl = fixHashCollision()
	//}

	_, err = ap.db.PostUrl(shortUrl, fullUrl)
	if err != nil {
		//ap.cfg.ErrorLog.Println("Failed to insert value in the database")
		return "", err
	}
	// cut prefix when search in db
	return ap.cfg.UrlHost + shortUrl, nil
}

func (ap *App) GetFullUrl(shortUrl string) (string, error) {
	fullUrl, err := ap.db.GetUrl(shortUrl[len(ap.cfg.UrlHost):])
	if err != nil {
		//ap.cfg.ErrorLog.Println("Failed to find value in the database")
		return "", err
	}
	return fullUrl, nil
}
