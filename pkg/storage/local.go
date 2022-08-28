package storage

import (
	"errors"
	"github.com/aytero/go-url-shortener-service/pkg/config"
	"github.com/aytero/go-url-shortener-service/pkg/model"
	"sync"
)

// need Mutex

type DatabaseLocal struct {
	db  []model.DbEntry
	cfg *config.Config
	mu  *sync.Mutex
}

func NewLocalStorage(cfg *config.Config) *DatabaseLocal {
	db, _ := InitLocalStorage()
	return &DatabaseLocal{
		db:  db,
		cfg: cfg,
		mu:  &sync.Mutex{},
	}
}

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
	return "", errors.New("not found")
}

func (engine *DatabaseLocal) GetShortUrl(fullUrl string) (string, error) {
	for _, a := range engine.db {
		if a.FullUrl == fullUrl {
			return a.ShortUrl, nil
		}
	}
	return "", errors.New("not found")
}

func (engine *DatabaseLocal) Lock() {
	engine.mu.Lock()
}

func (engine *DatabaseLocal) Unlock() {
	engine.mu.Unlock()
}
