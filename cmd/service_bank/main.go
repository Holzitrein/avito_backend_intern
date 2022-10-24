package main

import (
	"avito_balance/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.Routes(router)
	router.Run()
}
