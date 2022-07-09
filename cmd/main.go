package main

import (
	"github.com/aytero/ozon-fintech-url-service/internal/api"
	"github.com/aytero/ozon-fintech-url-service/internal/app"
	"github.com/aytero/ozon-fintech-url-service/pkg/config"
	"github.com/aytero/ozon-fintech-url-service/pkg/storage"
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
