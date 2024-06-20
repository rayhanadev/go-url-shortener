package models

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(url string) string {
    h := sha256.New()
    h.Write([]byte(url))
    hash := h.Sum(nil)
    slug := base64.URLEncoding.EncodeToString(hash)[:8]
    return slug
}
