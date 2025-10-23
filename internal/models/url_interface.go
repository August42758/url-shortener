package models

type UrlModelIntreface interface {
	AddUrl(shortUrl, originalUrl string) error
	GetOriginalUrl(shortUrl string) (string, error)
	IncreaseRedirectCount(shortUrl string) error
	GetUrlInfo(shortUrl string) (*Url, error)
}
