package app

import (
	"math/rand"
	"time"
	"urlShortener/pkg/config"
	"urlShortener/pkg/dataStorage"
)

type App struct {
	db  *dataStorage.Database
	cfg *config.Config
}

func NewApp(db *dataStorage.Database, cfg *config.Config) *App {
	rand.Seed(time.Now().UnixNano())
	return &App{
		db:  db,
		cfg: cfg,
	}
}
