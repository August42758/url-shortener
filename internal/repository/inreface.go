package repository

import (
	"urlShortener/internal/database"
)

type RepositoryShortenerIntreface interface {
	AddUrl(shortUrl, originalUrl string) error
	GetOriginalUrl(shortUrl string) (string, error)
	IncreaseRedirectCount(shortUrl string) error
	GetUrlInfo(shortUrl string) (database.UrlInfo, error)
}
