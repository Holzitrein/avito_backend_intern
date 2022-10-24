package database

import (
	"database/sql"
	"log"
	"time"
)

func Check(db *sql.DB) {
	time.Sleep(30 * time.Second)
	log.Println("rows")
	log.Println(db.Ping())
}
