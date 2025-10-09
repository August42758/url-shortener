package database

import (
	"database/sql"

	"urlShortener/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", GetDbAddres())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func GetDbAddres() string {
	addr := "postgres://"
	addr += config.AppConfig.DbConnect.Username + ":"
	addr += config.AppConfig.DbConnect.Password + "@"
	addr += config.AppConfig.DbConnect.Addres + "/"
	addr += config.AppConfig.DbConnect.Name + "?sslmode=disable"

	return addr
}
