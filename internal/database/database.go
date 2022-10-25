package database

import (
	"avito_balance/internal/model"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func ReserveBalanceDb(data model.ReserveCreate) model.ReserveCreateReturn {
	var returnData model.ReserveCreateReturn
	var isExists bool
	var isExists_2 bool
	var balance float32
	var reserve float32
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM users WHERE idUser = $1)", data.UserId).Scan(&isExists)
	if !(isExists) {
		returnData.Status = "User not found"
		return returnData
	}
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM services WHERE idService = $1)", data.IdService).Scan(&isExists_2)
	if !(isExists_2) {
		returnData.Status = "Service not found"
		return returnData
	}
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM orders WHERE idorder = $1)", data.IdOrder).Scan(&isExists_2)
	if isExists_2 {
		returnData.Status = "This order already exists"
		return returnData
	}
	db.QueryRow(context.Background(), "SELECT balance, reserve FROM users WHERE idUser = $1", data.UserId).Scan(&balance, &reserve)

	if (balance - reserve - data.Price) < 0 {
		returnData.Status = "not enough money"
		return returnData
	}
	tx, err := db.Begin(context.Background())
	if err != nil {
		returnData.Status = "Commit error"
		return returnData
	}
	defer tx.Rollback(context.Background())
	tx.Exec(context.Background(), "UPDATE users SET reserve=reserve+$1 WHERE idUser = $2", data.Price, data.UserId)
	response := "reserve"
	tx.Exec(context.Background(), "INSERT INTO orders (idorder,iduser,idservice,price,created,statusorder) VALUES ($1, $2, $3, $4, $5, $6)", data.IdOrder, data.UserId, data.IdService, data.Price, time.Now(), response)
	err = tx.Commit(context.Background())
	if err != nil {
		returnData.Status = "Commit error"
		return returnData
	}
	returnData.Status = "Successfully"
	return returnData
}

func ConfirmBalanceDb(data model.ReserveConfirm) model.ReserveCreateReturn {
	var returnData model.ReserveCreateReturn
	var isExists bool
	status := "reserve"
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM orders WHERE idorder = $1 AND iduser = $2 AND idservice = $3 AND price = $4 AND statusorder = $5)", data.IdOrder, data.UserId, data.IdService, data.Price, status).Scan(&isExists)
	if !(isExists) {
		returnData.Status = "Order not found"
		return returnData
	}
	if data.Command == "approved" {
		tx, err := db.Begin(context.Background())
		if err != nil {
			returnData.Status = "Commit error"
			return returnData
		}
		defer tx.Rollback(context.Background())
		tx.Exec(context.Background(), "UPDATE users SET reserve = reserve - $2 WHERE idUser = $1", data.UserId, data.Price)
		tx.Exec(context.Background(), "UPDATE users SET balance = balance - $2 WHERE idUser = $1", data.UserId, data.Price)
		response := "approved"
		tx.Exec(context.Background(), "UPDATE orders SET statusorder = $1 WHERE idorder = $2", response, data.IdOrder)
		err = tx.Commit(context.Background())
		if err != nil {
			returnData.Status = "Commit error"
			return returnData
		}
	}
	if data.Command == "cancel" {
		tx, err := db.Begin(context.Background())
		if err != nil {
			returnData.Status = "Commit error"
			return returnData
		}
		defer tx.Rollback(context.Background())
		tx.Exec(context.Background(), "UPDATE users SET reserve = reserve - $2 WHERE idUser = $1", data.UserId, data.Price)
		response := "cancel"
		tx.Exec(context.Background(), "UPDATE orders SET statusorder = $1 WHERE idorder = $2", response, data.IdOrder)
		err = tx.Commit(context.Background())
		if err != nil {
			returnData.Status = "Commit error"
			return returnData
		}
	}
	returnData.Status = "Successfully"
	return returnData
}

func TrunsferBalanceBd(data model.TransferBalanceUser) model.TransferBalanceUserReturn {
	var returnData model.TransferBalanceUserReturn
	var isExists bool
	var balance float32
	var reserve float32
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM users WHERE iduser = $1)", data.UserId).Scan(&isExists)
	if !(isExists) {
		returnData.Status = "User not found"
		return returnData
	}
	db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM users WHERE iduser = $1)", data.UserId2).Scan(&isExists)
	if !(isExists) {
		returnData.Status = "User not found"
		return returnData
	}
	db.QueryRow(context.Background(), "SELECT balance, reserve FROM users WHERE idUser = $1", data.UserId).Scan(&balance, &reserve)
	if (balance - reserve - data.Money) < 0 {
		returnData.Status = "not enough money"
		return returnData
	}

	tx, err := db.Begin(context.Background())
	if err != nil {
		returnData.Status = "Commit error"
		return returnData
	}
	defer tx.Rollback(context.Background())

	tx.Exec(context.Background(), "UPDATE users SET balance = balance - $1 WHERE idUser = $2", data.Money, data.UserId)
	tx.Exec(context.Background(), "UPDATE users SET balance = balance + $1 WHERE idUser = $2", data.Money, data.UserId2)
	err = tx.Commit(context.Background())
	if err != nil {
		returnData.Status = "Commit error"
		return returnData
	}
	returnData.Status = "Successfully"
	return returnData
}

func ReportServiceBd(data model.ReportServiceStruct) model.ReportServiceStructReturn {
	var returnData model.ReportServiceStructReturn
	year := strings.Split(data.Date, "-")[0]
	year_int, _ := strconv.Atoi(year)
	mounth := strings.Split(data.Date, "-")[1]
	mounth_int, _ := strconv.Atoi(mounth)
	date_from := time.Date(year_int, time.Month(mounth_int), 1, 0, 0, 0, 0, time.UTC)
	date_to := date_from.AddDate(0, 1, -1)

	rows, err := db.Query(context.Background(), "SELECT idservice, nameser FROM services")
	if err != nil {
		log.Println("Error find servises ")
		log.Println(err)
	}

	var idSer int
	var nameSer string
	var money float32
	data_final := [][]string{{"Услуга", "Доход"}}
	nameFile := "report_" + data.Date + ".csv"
	nameFile_os := "internal/database/csv/" + nameFile
	status := "approved"
	f, _ := os.Create(nameFile_os)
	defer f.Close()
	var num_rows = []model.ServiceName{}
	for rows.Next() {
		err := rows.Scan(&idSer, &nameSer)
		if err != nil {
			log.Fatal("1")
		}
		num_rows = append(num_rows, model.ServiceName{Id: idSer, Name: nameSer})

	}
	rows.Close()
	for i := 0; i < len(num_rows); i++ {
		rows_final, err := db.Query(context.Background(), "SELECT price FROM orders WHERE idService =  $1 AND statusOrder = $2 AND created > $3 AND created < $4",
			num_rows[i].Id, status, date_from, date_to)
		if err != nil {
			log.Println(err)
			log.Println("err")
		}
		var final_money float32
		final_money = 0
		for rows_final.Next() {
			err := rows_final.Scan(&money)
			log.Println(money)
			final_money = final_money + money
			if err != nil {
				log.Fatal("3")
			}
		}
		s := fmt.Sprintf("%v", final_money)
		data_temp := []string{num_rows[i].Name, s}
		data_final = append(data_final, data_temp)
		rows_final.Close()
	}

	w := csv.NewWriter(f)

	w.WriteAll(data_final)
	log.Print(data_final)
	returnData.Url = "localhost:8000/csv/" + nameFile
	return returnData
}
