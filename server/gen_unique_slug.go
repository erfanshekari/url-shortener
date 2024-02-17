package server

import (
	"math/rand"
	"time"

	"github.com/erfanshekari/url-shortener/models/link"
)

func genUniqueString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func genUniqueSlug(length int) string {
	slug := genUniqueString(length)
	l, err := link.FindBySlug(slug)
	if err != nil {
		panic(err)
	}
	if l != nil {
		return genUniqueSlug(length)
	}
	return "l" + slug
}
