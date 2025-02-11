package shortener

import (
	"1_increment_http_server/internal/providers/memory_provider"
	"math/rand"
)

var (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Shortener struct {
	mp memory_provider.MemoryProvider
}

func New(mp memory_provider.MemoryProvider) *Shortener {
	return &Shortener{
		mp: mp,
	}
}

func (s *Shortener) Shorten(url string) (string, error) {
	return "", nil
}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
