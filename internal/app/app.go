package app

import (
	"github.com/aytero/ozon-fintech-url-service/pkg/config"
	"github.com/aytero/ozon-fintech-url-service/pkg/storage"
	"math/rand"
	"time"
)

type App struct {
	db  *storage.Database
	cfg *config.Config
}

func NewApp(db *storage.Database, cfg *config.Config) *App {
	rand.Seed(time.Now().UnixNano())
	return &App{
		db:  db,
		cfg: cfg,
	}
}
