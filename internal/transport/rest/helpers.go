package rest

import (
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

var (
	shortUrlLenght int = 6
)

func (S *Shortener) genShortUrl() string {
	bytes := make([]byte, shortUrlLenght)

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i != shortUrlLenght; i++ {
		bytes[i] = byte(generator.Intn(255))
	}

	encoded := base64.URLEncoding.EncodeToString(bytes)
	encoded = strings.TrimRight(encoded, "=")
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")

	return encoded[:shortUrlLenght]
}
