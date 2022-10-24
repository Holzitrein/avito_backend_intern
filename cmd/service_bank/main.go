package main

import (
	"avito_balance/internal/database"
	"avito_balance/internal/routers"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

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

	router.Run()
}
