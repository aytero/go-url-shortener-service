package storage

import (
	"errors"
	"github.com/aytero/go-url-shortener-service/pkg/config"
)

type DatabaseMethodHandler interface {
	GetUrl(shortUrl string) (string, error)
	PostUrl(shortUrl string, fullUrl string) (string, error)
	GetShortUrl(fullUrl string) (string, error)
	Lock()
	Unlock()
}

type Database struct {
	cfg *config.Config
	DatabaseMethodHandler
}

func NewStorage(cfg *config.Config) (*Database, error) {
	switch cfg.StorageType {
	case "postgres":
		return &Database{
			cfg:                   cfg,
			DatabaseMethodHandler: NewPostgres(cfg),
		}, nil
	case "local":
		return &Database{
			cfg:                   cfg,
			DatabaseMethodHandler: NewLocalStorage(cfg),
		}, nil
	default:
		return &Database{
			cfg:                   cfg,
			DatabaseMethodHandler: NewLocalStorage(cfg),
		}, errors.New("Unknown data storage type. Expected: 'postgres' or 'local'. Usage: go run cmd/main.gp -storage=postgres")
	}
}
