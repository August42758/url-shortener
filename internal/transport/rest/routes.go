package rest

import (
	"net/http"
)

func (s *Shortener) SetRoutes() http.Handler {
	router := http.NewServeMux()

	router.Handle("/s/{short_url}", s.RedirectByCachedUrl(http.HandlerFunc(s.HandleRedirectByShortUrl)))
	router.HandleFunc("/shorten", s.HandleCreateShortUrl)
	router.HandleFunc("/analytycs/{short_url}", s.HandleGetUrlInfo)

	return router
}
