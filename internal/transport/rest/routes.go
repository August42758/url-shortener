package rest

import (
	"net/http"
)

func (s *Shortener) SetRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/s/{short_url}", s.HandleRedirectByShortUrl)
	router.HandleFunc("/shorten", s.HandleCreateShortUrl)
	router.HandleFunc("/analytycs/{short_url}", s.HandleGetUrlInfo)

	return router
}
