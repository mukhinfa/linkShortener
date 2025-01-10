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

// NewLink creates a new link with the given URL and unique hash
func NewLink(url string) *Link {
	return &Link{
		URL:  url,
		Hash: RandStringRunes(10),
	}
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
