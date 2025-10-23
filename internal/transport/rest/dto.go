package rest

import (
	"regexp"
)

var urlRegex = regexp.MustCompile("^(https?|ftp)://[^\\s/$.?#].[^\\s]*$")

// request dto
type OriginalUrlDTO struct {
	Url string `json:"url"`
}

func (o OriginalUrlDTO) Validate() bool {
	return urlRegex.Match([]byte(o.Url))
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
type UrlInfoDTO struct {
	ShortUrl      string `json:"short_url"`
	OriginalUrl   string `json:"original_url"`
	RedirectCount int    `json: "redirect_count"`
}

func NewUrlInfoDto(shortUrl string, originalUrl string, redirectCount int) *UrlInfoDTO {
	return &UrlInfoDTO{
		ShortUrl:      shortUrl,
		OriginalUrl:   originalUrl,
		RedirectCount: redirectCount,
	}
}
