package main

import (
	"rayhanadev/go-shrt/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
