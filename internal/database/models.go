package database

type UrlInfo struct {
	Id            int
	ShortUrl      string
	OriginalUrl   string
	RedirectCount int
}

func NEwUrlInfo() UrlInfo {
	return UrlInfo{}
}
