package database

import "avito_balance/internal/model"

func ReportServiceBd(data model.ReportServiceStruct) model.ReportServiceStructReturn {
	var returnData model.ReportServiceStructReturn
	err := db.QueryRow(context.Background(), "SELECT EXISTS (SELECT * FROM users WHERE idUser = $1)", data.UserId).Scan(&isExists)
	if err != nil {
		returnData.Status = "False"
		return returnData
	}
	return returnData
}
