package models

type Url struct {
	Id            int
	ShortUrl      string
	OriginalUrl   string
	RedirectCount int
}

func NewUrl() *Url {
	return &Url{}
}
