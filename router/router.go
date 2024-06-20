package router

import (
	"rayhanadev/go-shrt/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/", handlers.HomeHandler)
    r.POST("/shorten", handlers.ShortenURLHandler)
    r.GET("/:slug", handlers.RedirectHandler)
    return r
}
