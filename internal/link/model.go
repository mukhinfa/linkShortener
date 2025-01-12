package link

import (
	"math/rand"

	"gorm.io/gorm"
)

// Link represents a link with a URL and a unique hash
type Link struct {
	gorm.Model
	URL  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

// NewLink creates a new link with the given URL
func NewLink(url string) *Link {
	return &Link{
		URL: url,
	}
}

// GenerateHash generates a unique hash for the link
func (l *Link) GenerateHash(handler *Handler) string {
	for {
		hash := RandStringRunes(10)
		if _, err := handler.Repository.GetByHash(hash); err != nil {
			l.Hash = hash
			break
		}
	}

	return l.Hash
}

var letterRunes = []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ")

// RandStringRunes generates a random string of a given length from alphabetic characters
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
