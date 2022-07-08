package app

import (
	"urlShortener/pkg/config"
	"urlShortener/pkg/dataStorage"
)

type App struct {
	db  *dataStorage.Database
	cfg *config.Config
}

func NewApp(db *dataStorage.Database, cfg *config.Config) *App {
	return &App{
		db:  db,
		cfg: cfg,
	}
}
