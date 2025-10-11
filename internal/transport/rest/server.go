package rest

import (
	"errors"
	"net/http"

	"urlShortener/internal/config"

	"github.com/gorilla/mux"
)

type HttpServerShortener struct {
	HttpHandlers HttpHandlersShortener
}

func NewHttpServerShortener(httpHandlers HttpHandlersShortener) HttpServerShortener {
	return HttpServerShortener{
		HttpHandlers: httpHandlers,
	}
}

func (h HttpServerShortener) Start() error {
	router := mux.NewRouter()

	router.Path("/s/{short_url}").Methods("GET").HandlerFunc(h.HttpHandlers.HandleRedirectByShortUrl)
	router.Path("/shorten").Methods("POST").HandlerFunc(h.HttpHandlers.HandleCreateShortUrl)
	router.Path("/analytycs/{short_url}").Methods("GET").HandlerFunc(h.HttpHandlers.HandleGetUrlInfo)

	err := http.ListenAndServe(config.AppConfig.ServerAddres, router)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
