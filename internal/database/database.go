package database

import (
	"avito_balance/internal/model"
	"context"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func SetDb(db_new *pgx.Conn) {
	db = db_new
}

func GetBalanceDb(data model.BalanceGet) model.BalanceGetReturn {
	var money float32
	var returnData model.BalanceGetReturn
	err := db.QueryRow(context.Background(), "select balance from Users where idUser=$1", data.UserId).Scan(&money)
	if err != nil {
		panic("getBalance Erorr")
	}
	returnData.Balance = money
	return returnData
}
