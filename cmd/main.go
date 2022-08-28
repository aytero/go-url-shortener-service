package main

import (
	"github.com/aytero/go-url-shortener-service/internal/api"
	"github.com/aytero/go-url-shortener-service/internal/app"
	"github.com/aytero/go-url-shortener-service/pkg/config"
	"github.com/aytero/go-url-shortener-service/pkg/storage"
)

func main() {
	cfg := config.NewConfig(".env")
	db, err := storage.NewStorage(cfg)
	if err != nil {
		cfg.ErrorLog.Fatal(err)
	}
	newApp := app.NewApp(db, cfg)
	newApi := api.NewApi(newApp, cfg)
	router := newApi.NewRouter()

	cfg.InfoLog.Println("Config finished. Starting server.")
	cfg.ErrorLog.Fatal(router.Run(":8080"))
}
