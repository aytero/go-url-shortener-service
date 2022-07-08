package main

import (
	"urlShortener/internal/api"
	"urlShortener/internal/app"
	"urlShortener/pkg/config"
	"urlShortener/pkg/dataStorage"
)

func main() {

	cfg := config.NewConfig(".env")
	db := dataStorage.NewDataStorage(cfg)
	newApp := app.NewApp(db, cfg)
	newApi := api.NewApi(newApp, cfg)
	router := newApi.NewRouter()

	cfg.InfoLog.Println("Config finished. Starting server.")
	cfg.ErrorLog.Fatal(router.Run(":8080"))

	//server := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        mux,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
}
