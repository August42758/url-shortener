package rest

import (
	"log"
	"urlShortener/internal/models"

	"github.com/redis/go-redis/v9"
)

type Shortener struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	UrlModel    models.UrlModelIntreface
	Redis       *redis.Client
}

func NewShortener(redis *redis.Client, urlModel models.UrlModelIntreface, errorLogger, InfoLogger *log.Logger) *Shortener {
	return &Shortener{
		ErrorLogger: errorLogger,
		InfoLogger:  InfoLogger,
		UrlModel:    urlModel,
		Redis:       redis,
	}
}
