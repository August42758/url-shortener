package rest

import (
	"encoding/json"
	"time"
)

// request dto
type OriginalUrlDTO struct {
	Url string `json:"url"`
}

func (o OriginalUrlDTO) Validate() error {
	if o.Url == "" {
		return errWrongJsonFieldValue
	}
	return nil
}

// response dto
type ShortUrlDTO struct {
	ShortUrl string `json:"short_url"`
}

func NewShortUrlDTO(shortUrl string) ShortUrlDTO {
	return ShortUrlDTO{
		ShortUrl: shortUrl,
	}
}

// response dto
type ErrorDTO struct {
	Err     string    `json:"error"`
	ErrTime time.Time `json:"error_time"`
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Err:     err.Error(),
		ErrTime: time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// response dto
type UrlInfoDTO struct {
	ShortUrl      string `json:"short_url"`
	OriginalUrl   string `json:"original_url"`
	RedirectCount int    `json: "redirect_count"`
}

func NewUrlInfoDto(shortUrl string, originalUrl string, redirectCount int) UrlInfoDTO {
	return UrlInfoDTO{
		ShortUrl:      shortUrl,
		OriginalUrl:   originalUrl,
		RedirectCount: redirectCount,
	}
}
