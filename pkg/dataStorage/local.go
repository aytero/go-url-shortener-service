package dataStorage

import (
	"urlShortener/pkg/config"
	"urlShortener/pkg/model"
)

// need Mutex

type DatabaseLocal struct {
	db  []model.DbEntry
	cfg *config.Config
}

func NewLocalStorage(cfg *config.Config) *DatabaseLocal {
	db, _ := InitLocalStorage()
	return &DatabaseLocal{
		db:  db,
		cfg: cfg,
	}
}

//{ShortUrl: "aaa", FullUrl: "example.com/aaafds"},
//{ShortUrl: "aab", FullUrl: "example.com/aaetccsa"},

func InitLocalStorage() ([]model.DbEntry, error) {
	return []model.DbEntry{}, nil
}

func (engine *DatabaseLocal) PostUrl(shortUrl string, fullUrl string) (string, error) {
	newEntry := model.DbEntry{ShortUrl: shortUrl, FullUrl: fullUrl}
	engine.db = append(engine.db, newEntry)
	return shortUrl, nil
}

func (engine *DatabaseLocal) GetUrl(shortUrl string) (string, error) {
	for _, a := range engine.db {
		if a.ShortUrl == shortUrl {
			return a.FullUrl, nil
		}
	}
	return "", nil
}
