package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"urlShortener/internal/service"

	"github.com/gorilla/mux"
)

type HttpHandlersShortener struct {
	service service.ServiceShortener
}

func NewHttpHandlersShortener(service service.ServiceShortener) HttpHandlersShortener {
	return HttpHandlersShortener{
		service: service,
	}
}

func (h *HttpHandlersShortener) HandleCreateShortUrl(w http.ResponseWriter, r *http.Request) {
	defer catchPanic("HandleCreateShortUrl")

	httpRequestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
		return
	}

	var originalUrlDTO OriginalUrlDTO
	if err := json.Unmarshal(httpRequestBytes, &originalUrlDTO); err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err = originalUrlDTO.Validate(); err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	shortUrl := h.service.AddOriginalUrl(originalUrlDTO.Url)

	shortUrlDTO := NewShortUrlDTO(shortUrl)
	b, err := json.Marshal(shortUrlDTO)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (h HttpHandlersShortener) HandleRedirectByShortUrl(w http.ResponseWriter, r *http.Request) {
	defer catchPanic("HandlerRedirectByShortUrl")

	shortUrl := mux.Vars(r)["short_url"]

	originalUrl, err := h.service.GetOriginalUrl(shortUrl)
	if err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusNotFound)
		return
	}

	if err := h.service.IncreaseRedirectCount(shortUrl); err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}

func (h HttpHandlersShortener) HandleGetUrlInfo(w http.ResponseWriter, r *http.Request) {
	defer catchPanic("HandnleGetUrlInfo")

	shortUrl := mux.Vars(r)["short_url"]

	urlInfoModel, err := h.service.GetUrlInfo(shortUrl)
	if err != nil {
		errorDTO := NewErrorDTO(err)
		http.Error(w, errorDTO.ToString(), http.StatusNotFound)
		return
	}

	urlInfoDTO := NewUrlInfoDto(urlInfoModel.ShortUrl, urlInfoModel.OriginalUrl, urlInfoModel.RedirectCount)

	b, err := json.Marshal(&urlInfoDTO)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
