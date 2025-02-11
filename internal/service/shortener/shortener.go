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

// ShortenJoin - функция для объединения полного url с коротким.
func (s *Shortener) ShortenJoin(url string) {
	shortenLink := randStr(6)
	s.mp.SetValueMemory(shortenLink, url)
}

// randStr - функция генерации случайного короткого адреса.
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// Случайно выбирает символ из charset.
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
