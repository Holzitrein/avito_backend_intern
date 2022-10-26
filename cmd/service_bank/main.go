package main

import (
	"avito_balance/internal/database"
	"avito_balance/internal/routers"
	"context"
	"os"

	_ "avito_balance/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title       REST server for balance
// @version     1.0
// @description Task from avito for backend intern

// @contact.name Andreychuk Andrew
// @contact.url  https://vk.com/holzitrein

// @host     localhost:8000
// @BasePath /

func main() {
	router := gin.Default()
	routers.Routes(router)

	info_db := "postgres://avito:avito@database:5432/avito"
	conn, err := pgx.Connect(context.Background(), info_db)
	if err != nil {
		os.Exit(1)
	}
	database.SetDb(conn)

	defer conn.Close(context.Background())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
