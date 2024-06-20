package main

import (
	"rayhanadev/url-shortener/database"
	"rayhanadev/url-shortener/router"
)

func main() {
    database.InitDB()
    r := router.SetupRouter()
    r.Run(":8080")
}
