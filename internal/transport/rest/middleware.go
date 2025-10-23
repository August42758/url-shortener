package rest

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func (s *Shortener) RedirectByCachedUrl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, err := s.Redis.Get(context.Background(), r.PathValue("short_url")).Result()
		if err == redis.Nil {
			next.ServeHTTP(w, r)
			return
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, val, http.StatusTemporaryRedirect)
	})
}
