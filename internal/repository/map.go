package repository

type RepositoryShortenerMap struct {
	storage map[string]string
}

func NewRepositoryShortenerMap() RepositoryShortenerMap {
	return RepositoryShortenerMap{
		storage: map[string]string{},
	}
}

func (u *RepositoryShortenerMap) AddUrl(shortUrl, originalUrl string) error {
	if _, ok := u.storage[shortUrl]; ok {
		return ErrUrlAlreadyExists
	}

	u.storage[shortUrl] = originalUrl

	return nil
}

func (u *RepositoryShortenerMap) GetOriginalUrl(shortUrl string) (string, error) {
	if _, ok := u.storage[shortUrl]; !ok {
		return "", ErrUrlNotFound
	}

	return u.storage[shortUrl], nil
}
