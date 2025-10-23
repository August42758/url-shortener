package models

type UrlModelMap struct {
	storage map[string]string
}

func NewUrlModelMap() *UrlModelMap {
	return &UrlModelMap{
		storage: map[string]string{},
	}
}

func (u *UrlModelMap) AddUrl(shortUrl, originalUrl string) error {
	if _, ok := u.storage[shortUrl]; ok {
		return ErrUrlAlreadyExists
	}

	u.storage[shortUrl] = originalUrl

	return nil
}

func (u *UrlModelMap) GetOriginalUrl(shortUrl string) (string, error) {
	if _, ok := u.storage[shortUrl]; !ok {
		return "", ErrUrlNotFound
	}

	return u.storage[shortUrl], nil
}
