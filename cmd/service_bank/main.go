package main

import (
	"avito_balance/internal/routers"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	routers.Routes(router)

	info_db := "postgres://avito:avito@database:5432/avito"
	conn, err := pgx.Connect(context.Background(), info_db)
	if err != nil {
		log.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var some string
	var some2 string

	err = conn.QueryRow(context.Background(), "select balance, reserve from Users where idUser=$1", "2").Scan(&some, &some2)
	if err != nil {
		log.Print(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	log.Println(some, " ", some2)
	router.Run()
}
