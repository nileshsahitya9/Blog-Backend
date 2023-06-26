package db

import (
	"database/sql"
	"fmt"

	"github.com/nileshsahitya9/Blogs-Backend/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func PostgreSQLDB() (*sql.DB, error) {
	cfg, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase))
	if err != nil {

		return nil, err
	}
	fmt.Println(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&client_encoding=utf8", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase))

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}
