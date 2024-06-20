package storage

import (
	"rayhanadev/url-shortener/database"
	"rayhanadev/url-shortener/models"
)

func SaveURL(url models.URL) error {
    return database.DB.Create(&url).Error
}

func GetURL(shortURL string) (models.URL, bool) {
    var url models.URL
    result := database.DB.Where("short_url = ? OR custom_url = ?", shortURL, shortURL).First(&url)
    if result.Error != nil {
        return url, false
    }
    return url, true
}
