package main

import (
	"avito_balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Hello)
	router.Run()
}
