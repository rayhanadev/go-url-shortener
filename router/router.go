package router

import (
	"rayhanadev/url-shortener/handlers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

    r.GET("/", handlers.HomeHandler)
    r.POST("/register", handlers.RegisterHandler)
    r.POST("/login", handlers.LoginHandler)
    r.POST("/logout", handlers.LogoutHandler)
    r.POST("/shorten", handlers.ShortenURLHandler)
    r.GET("/:slug", handlers.RedirectHandler)

    return r
}
