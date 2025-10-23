package rest

import (
	"log"
	"urlShortener/internal/models"
)

type Shortener struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	UrlModel    models.UrlModelIntreface
}

func NewShortener(urlModel models.UrlModelIntreface, errorLogger, InfoLogger *log.Logger) *Shortener {
	return &Shortener{
		ErrorLogger: errorLogger,
		InfoLogger:  InfoLogger,
		UrlModel:    urlModel,
	}
}
