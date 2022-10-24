package database

import (
	"avito_balance/internal/model"
	"context"
	"time"

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
		//panic("getBalance Erorr")
	}
	returnData.Balance = money
	return returnData
}

func AddBalanceDb(data model.BalanceAdd) model.BalanceAddReturn {
	var isExists bool
	var returnData model.BalanceAddReturn
	err := db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM users WHERE idUser = $1)", data.UserId).Scan(&isExists)
	if err != nil {
		returnData.Status = "False"
		return returnData
	}
	if isExists {
		db.Exec(context.Background(), "UPDATE users SET balance=$2 + balance WHERE idUser = $1", data.UserId, data.Money)

	} else {
		db.Exec(context.Background(), "INSERT INTO users (idUser ,balance,reserve) VALUES ( $1, $2, 0); ", data.UserId, data.Money)
	}
	db.Exec(context.Background(), "INSERT INTO history_add (amount , idUser, timeAdd, note) VALUES ( $1, $2, $3, $4); ",
		data.Money, data.UserId, time.Now(), data.Note)
	returnData.Status = "true"
	return returnData
}
