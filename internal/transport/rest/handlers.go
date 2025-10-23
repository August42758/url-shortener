package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	"urlShortener/internal/models"
)

func (s *Shortener) HandleCreateShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	httpRequestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.ErrorLogger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var originalUrlDTO OriginalUrlDTO
	if err := json.Unmarshal(httpRequestBytes, &originalUrlDTO); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !originalUrlDTO.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	shortUrl := s.genShortUrl()

	if err := s.UrlModel.AddUrl(shortUrl, originalUrlDTO.Url); err != nil {
		if errors.Is(err, models.ErrUrlAlreadyExists) {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		} else {
			s.ErrorLogger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	shortUrlDTO := NewShortUrlDTO(shortUrl)
	b, err := json.Marshal(shortUrlDTO)
	if err != nil {
		s.ErrorLogger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (s *Shortener) HandleRedirectByShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	shortUrl := r.PathValue("short_url")

	originalUrl, err := s.UrlModel.GetOriginalUrl(shortUrl)
	if err != nil {
		if errors.Is(err, models.ErrUrlNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			s.ErrorLogger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if err := s.UrlModel.IncreaseRedirectCount(shortUrl); err != nil {
		s.ErrorLogger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := s.Redis.Set(context.Background(), shortUrl, originalUrl, 30*time.Minute).Err(); err != nil {
		s.ErrorLogger.Println(err)
	}

	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}

func (s *Shortener) HandleGetUrlInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	shortUrl := r.PathValue("short_url")

	urlInfoModel, err := s.UrlModel.GetUrlInfo(shortUrl)
	if err != nil {
		if errors.Is(err, models.ErrUrlNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			s.ErrorLogger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	urlInfoDTO := NewUrlInfoDto(urlInfoModel.ShortUrl, urlInfoModel.OriginalUrl, urlInfoModel.RedirectCount)

	b, err := json.Marshal(urlInfoDTO)
	if err != nil {
		s.ErrorLogger.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
