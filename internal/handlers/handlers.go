package handlers

import (
	"avito_balance/internal/database"
	"avito_balance/internal/model"
	"avito_balance/internal/validate"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) { //Получение баланса
	var requestBody model.BalanceGet
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	reply := database.GetBalanceDb(requestBody)
	c.JSON(200, reply)
}

func TransferBalance(c *gin.Context) {
	var requestBody model.TransferBalanceUser
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.TransferBalanceValidate(requestBody) {
		reply := database.TrunsferBalanceBd(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}

func AddBalance(c *gin.Context) {
	var requestBody model.BalanceAdd
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.AddBalanceValidate(requestBody) {
		reply := database.AddBalanceDb(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}

func ReserveBalance(c *gin.Context) {
	var requestBody model.ReserveCreate
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.ReserveBalanceValidate(requestBody) {
		reply := database.ReserveBalanceDb(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}

func ConformBalnce(c *gin.Context) {
	var requestBody model.ReserveConfirm
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.ConformBalnceValidate(requestBody) {
		reply := database.ConfirmBalanceDb(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}

func ReportService(c *gin.Context) {
	var requestBody model.ReportServiceStruct
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.ReportServiceValidate(requestBody) {
		reply := database.ReportServiceBd(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}

func ReportServiceFile(c *gin.Context) {
	url := c.Param("file")
	c.File("internal/database/csv/" + url)
}

func ReportOperation(c *gin.Context) {
	var requestBody model.ReportOperationRequest
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	if validate.ReportOperationValidate(requestBody) {
		reply := database.ReportOperationBd(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}
