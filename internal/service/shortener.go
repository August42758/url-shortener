package service

import (
	"math/rand"

	"urlShortener/internal/repository"
)

type ServiceShortener struct {
	repository repository.RepositoryShortener
}

func NewServiceShortener(repository repository.RepositoryShortener) ServiceShortener {
	return ServiceShortener{
		repository: repository,
	}
}

func (s *ServiceShortener) AddOriginalUrl(originalUrl string) string {
	for {
		shortUrl := s.genShortUrl()
		if err := s.repository.AddUrl(shortUrl, originalUrl); err == nil {
			return shortUrl
		}
	}
}

func (s *ServiceShortener) GetOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := s.repository.GetOriginalUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (s ServiceShortener) genShortUrl() string {
	shortUrl := ""
	for i := 0; i != shortUrlLenght; i++ {
		shortUrl += string(availableSymbols[rand.Intn(len(availableSymbols))])
	}
	return shortUrl
}
