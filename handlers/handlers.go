package handlers

import (
	"net/http"
	"rayhanadev/url-shortener/database"
	"rayhanadev/url-shortener/models"
	"rayhanadev/url-shortener/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
    c.String(http.StatusOK, "Welcome to the URL Shortener!")
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{Username: req.Username}
    if err := user.SetPassword(req.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func LoginHandler(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    if !user.CheckPassword(req.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    session := sessions.Default(c)
    session.Set("user_id", user.ID)
    session.Save()

    c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func LogoutHandler(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

type ShortenRequest struct {
    URL       string `json:"url" binding:"required"`
    CustomURL string `json:"custom_url"`
}

type ShortenResponse struct {
    ShortURL string `json:"short_url"`
}

func ShortenURLHandler(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }

    var shortURL string
    if req.CustomURL != "" {
        shortURL = req.CustomURL
    } else {
        shortURL = models.GenerateShortURL(req.URL)
    }

    url := models.URL{
        OriginalURL: req.URL,
        ShortURL:    shortURL,
        CustomURL:   req.CustomURL,
        UserID:      userID.(uint),
    }

    if err := storage.SaveURL(url); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    resp := ShortenResponse{ShortURL: shortURL}
    c.JSON(http.StatusOK, resp)
}

func RedirectHandler(c *gin.Context) {
    shortURL := c.Param("shortURL")
    url, found := storage.GetURL(shortURL)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }
    c.Redirect(http.StatusFound, url.OriginalURL)
}
