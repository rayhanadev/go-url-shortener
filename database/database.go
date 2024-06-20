package database

import "sync"

var (
    urlStore = make(map[string]string)
    mu       sync.RWMutex
)

func SaveURL(slug, url string) {
    mu.Lock()
    defer mu.Unlock()
    urlStore[slug] = url
}

func GetURL(slug string) (string, bool) {
    mu.RLock()
    defer mu.RUnlock()
    url, found := urlStore[slug]
    return url, found
}
