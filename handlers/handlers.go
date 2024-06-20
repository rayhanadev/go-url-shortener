package handlers

import (
	"net/http"
	"rayhanadev/go-shrt/database"
	"rayhanadev/go-shrt/models"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
    URL string `json:"url" binding:"required"`
}

type ShortenResponse struct {
    ShortURL string `json:"slug"`
}

func HomeHandler(c *gin.Context) {
    c.String(http.StatusOK, "Welcome to the URL Shortener!")
}

func ShortenURLHandler(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    slug := models.GenerateShortURL(req.URL)
    database.SaveURL(slug, req.URL)

    resp := ShortenResponse{ShortURL: slug}
    c.JSON(http.StatusOK, resp)
}

func RedirectHandler(c *gin.Context) {
    slug := c.Param("slug")
    url, found := database.GetURL(slug)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }
    c.Redirect(http.StatusFound, url)
}
