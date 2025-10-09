package repository

type RepositoryShortenerIntreface interface {
	AddUrl(shortUrl, originalUrl string) error
	GetOriginalUrl(shortUrl string) (string, error)
}
