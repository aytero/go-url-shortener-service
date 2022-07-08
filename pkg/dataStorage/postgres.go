package dataStorage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"urlShortener/pkg/config"
)

type DatabasePostgres struct {
	db  *sql.DB
	cfg *config.Config
}

func NewPostgres(cfg *config.Config) *DatabasePostgres {
	conn, _ := InitPostgres(cfg)
	return &DatabasePostgres{
		db:  conn,
		cfg: cfg,
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
		engine.cfg.ErrorLog.Println(err)
	}
	//lastId, err := res.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		engine.cfg.ErrorLog.Println(err)
	}
	//log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	engine.cfg.InfoLog.Printf("affected = %d\n", rowCnt)
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
