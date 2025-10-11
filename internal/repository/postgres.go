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

func (r RepositoryShortenerPostgres) IncreaseRedirectCount(shortUrl string) error {
	var redirectCount int
	if err := r.db.QueryRow("SELECT redirect_count FROM urls WHERE short_url = $1", shortUrl).Scan(&redirectCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUrlNotFound
		} else {
			panic(err)
		}
	}

	redirectCount += 1
	r.db.Exec("UPDATE urls SET redirect_count = $1 WHERE short_url = $2", redirectCount, shortUrl)

	return nil
}

func (r RepositoryShortenerPostgres) GetUrlInfo(shortUrl string) (database.UrlInfo, error) {
	urlInfo := database.NEwUrlInfo()
	if err := r.db.QueryRow("SELECT * FROM urls WHERE short_url = $1", shortUrl).Scan(&urlInfo.Id, &urlInfo.ShortUrl, &urlInfo.OriginalUrl, &urlInfo.RedirectCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.UrlInfo{}, ErrUrlNotFound
		} else {
			panic(err)
		}
	}
	return urlInfo, nil
}
