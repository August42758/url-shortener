package models

import (
	"database/sql"
	"errors"
)

type UrlModelPostgres struct {
	db *sql.DB
}

func NewUrlModelPostgres(db *sql.DB) *UrlModelPostgres {
	return &UrlModelPostgres{
		db: db,
	}
}

func (r *UrlModelPostgres) AddUrl(shortUrl, originalUrl string) error {
	err := r.db.QueryRow("SELECT short_url FROM urls WHERE short_url = $1", shortUrl).Scan()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	_, err = r.db.Exec("INSERT INTO urls (short_url, original_url) VALUES($1, $2)", shortUrl, originalUrl)
	if err != nil {
		return err
	}

	return nil
}

func (r *UrlModelPostgres) GetOriginalUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := r.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortUrl).Scan(&originalUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUrlNotFound
		}
		return "", err
	}

	return originalUrl, nil
}

func (r *UrlModelPostgres) IncreaseRedirectCount(shortUrl string) error {
	var redirectCount int
	if err := r.db.QueryRow("SELECT redirect_count FROM urls WHERE short_url = $1", shortUrl).Scan(&redirectCount); err != nil {
		return err
	}

	redirectCount += 1
	if _, err := r.db.Exec("UPDATE urls SET redirect_count = $1 WHERE short_url = $2", redirectCount, shortUrl); err != nil {
		return err
	}

	return nil
}

func (r *UrlModelPostgres) GetUrlInfo(shortUrl string) (*Url, error) {
	urlInfo := NewUrl()
	if err := r.db.QueryRow("SELECT * FROM urls WHERE short_url = $1", shortUrl).Scan(&urlInfo.Id, &urlInfo.ShortUrl, &urlInfo.OriginalUrl, &urlInfo.RedirectCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUrlNotFound
		} else {
			return nil, err
		}
	}
	return urlInfo, nil
}
