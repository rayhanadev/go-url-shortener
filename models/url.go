package models

import (
	"crypto/sha256"
	"encoding/base64"

	"gorm.io/gorm"
)

type URL struct {
    gorm.Model
    OriginalURL string `gorm:"unique"`
    ShortURL    string `gorm:"unique"`
    CustomURL   string `gorm:"unique"`
    UserID      uint
}

func GenerateShortURL(url string) string {
    h := sha256.New()
    h.Write([]byte(url))
    hash := h.Sum(nil)
    shortURL := base64.URLEncoding.EncodeToString(hash)[:8]
    return shortURL
}
