package main

import (
	"one-go/router"
)

func main() {
	server := router.InitRouter()
	server.GinEngine.Run(":8000")
}
