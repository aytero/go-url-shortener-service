package dataStorage

import (
	"urlShortener/pkg/config"
)

type DatabaseMethodHandler interface {
	GetUrl(shortUrl string) (string, error)
	PostUrl(shortUrl string, fullUrl string) (string, error)
}

type Database struct {
	cfg *config.Config
	DatabaseMethodHandler
}

func NewDataStorage(cfg *config.Config) *Database {
	switch cfg.StorageType {
	case "postgres":
		return &Database{
			cfg:                   cfg,
			DatabaseMethodHandler: NewPostgres(cfg),
		}
	case "local":
		return &Database{
			cfg:                   cfg,
			DatabaseMethodHandler: NewLocalStorage(cfg),
		}
	default:
		cfg.ErrorLog.Fatalf(
			"Unknown data storage type. Expected: 'postgres' or 'local'. Usage: go run cmd/main.gp -storage=postgres")
	}
	return nil
}
