package dataStorage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
	"urlShortener/pkg/config"
)

type DatabasePostgres struct {
	db  *sql.DB
	cfg *config.Config
	mu  *sync.Mutex
}

func NewPostgres(cfg *config.Config) *DatabasePostgres {
	conn, _ := InitPostgres(cfg)
	return &DatabasePostgres{
		db:  conn,
		cfg: cfg,
		mu:  &sync.Mutex{},
	}
}

func InitPostgres(cfg *config.Config) (*sql.DB, error) {
	conn, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		cfg.ErrorLog.Fatalf("Unable to connect to database: %v\n", err)
	}
	//defer conn.Close()

	//err = conn.Ping()
	//if err != nil {
	//	return nil, err
	//}
	return conn, nil
}

func (engine *DatabasePostgres) PostUrl(shortUrl string, fullUrl string) (string, error) {
	stmt := "INSERT INTO urls (short_url, full_url) VALUES ($1, $2)"
	res, err := engine.db.Exec(stmt, shortUrl, fullUrl)
	if err != nil {
		return "", err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		engine.cfg.ErrorLog.Println(err)
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		engine.cfg.ErrorLog.Println(err)
		return "", err
	}
	engine.cfg.InfoLog.Printf("Id = %d, affected = %d\n", lastId, rowCnt)
	return shortUrl, nil
}

func (engine *DatabasePostgres) GetUrl(shortUrl string) (string, error) {
	var fullUrl string
	stmt := "SELECT full_url FROM urls WHERE short_url=$1"
	err := engine.db.QueryRow(stmt, shortUrl).Scan(&fullUrl)
	if err != nil {
		return "", err
	}
	return fullUrl, nil
}

func (engine *DatabasePostgres) GetShortUrl(fullUrl string) (string, error) {
	var shortUrl string
	stmt := "SELECT short_url FROM urls WHERE full_url=$1"
	err := engine.db.QueryRow(stmt, fullUrl).Scan(&shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (engine *DatabasePostgres) Lock() {
	engine.mu.Lock()
}

func (engine *DatabasePostgres) Unlock() {
	engine.mu.Unlock()
}
