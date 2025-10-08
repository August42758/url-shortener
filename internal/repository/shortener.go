package repository

type RepositoryShortener struct {
	storage map[string]string
}

func NewRepositoryShortener() RepositoryShortener {
	return RepositoryShortener{
		storage: map[string]string{},
	}
}

func (u *RepositoryShortener) AddUrl(shortUrl, originalUrl string) error {
	if _, ok := u.storage[shortUrl]; ok {
		return ErrUrlAlreadyExists
	}

	u.storage[shortUrl] = originalUrl

	return nil
}

func (u *RepositoryShortener) GetOriginalUrl(shortUrl string) (string, error) {
	if _, ok := u.storage[shortUrl]; !ok {
		return "", ErrUrlNotFound
	}

	return u.storage[shortUrl], nil
}
