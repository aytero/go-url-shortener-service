package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
	StorageType string
	UrlHost     string
	DatabaseUrl string
}

func NewConfig(envFile string) *Config {

	err := godotenv.Load(envFile)
	if err != nil {
		panic("Failed loading environment file")
	}
	cfg := &Config{
		InfoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		StorageType: getDbType(),
		UrlHost:     os.Getenv("URL_HOST"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
	return cfg
}

func getDbType() string {
	var dbFromFlag string

	flag.StringVar(&dbFromFlag, "storage", "local", "type of storage: postgres or local")
	flag.Parse()
	return dbFromFlag
}
