package repository

import (
	"database/sql"
	"errors"

	"urlShortener/internal/database"
)

type RepositoryShortenerPostgres struct {
	db *sql.DB
}

func NewRepositoryShortenerPostgres() RepositoryShortenerPostgres {
	return RepositoryShortenerPostgres{
		db: database.ConnectDB(),
	}
}

func (r RepositoryShortenerPostgres) AddUrl(shortUrl, originalUrl string) error {
	err := r.db.QueryRow("SELECT short_url FROM urls WHERE short_url = $1", shortUrl).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		} else {
			panic(err)
		}
	}

	_, err = r.db.Exec("INSERT INTO urls (short_url, original_url) VALUES($1, $2)", shortUrl, originalUrl)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r RepositoryShortenerPostgres) GetOriginalUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := r.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortUrl).Scan(&originalUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUrlNotFound
		}
		panic(err)
	}

	return originalUrl, nil
}
