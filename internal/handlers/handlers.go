package handlers

import (
	"avito_balance/internal/database"
	"avito_balance/internal/model"
	"avito_balance/internal/validate"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary     Get balance
// @Description get balance user by id
// @Tags        operation for user balance
// @Accept      json
// @Produce     json
// @Param       json model.BalanceGet true "json struct for request"
// @Success     200  {object}         model.BalanceGetReturn
// @Failure     400  {object}         model.ErrorReq
// @Router      /user [get]

func GetBalance(c *gin.Context) {
	var requestBody model.BalanceGet
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, "error")
	}
	reply := database.GetBalanceDb(requestBody)
	c.JSON(200, reply)
}

func TransferBalance(c *gin.Context) {
	var requestBody model.TransferBalanceUser
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, "error")
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
		c.JSON(400, "error")
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
		c.JSON(400, "error")
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
		c.JSON(400, "error")
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
		c.JSON(400, "error")
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
	if _, err := os.Stat("internal/database/csv/" + url); os.IsNotExist(err) {
		c.JSON(404, "File not gound")
	}
	c.File("internal/database/csv/" + url)
}

func ReportOperation(c *gin.Context) {
	var requestBody model.ReportOperationRequest
	var erMes model.ErrorReq
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, "error")
	}
	if validate.ReportOperationValidate(requestBody) {
		reply := database.ReportOperationBd(requestBody)
		c.JSON(200, reply)
	} else {
		erMes.ErrorMes = "Incorrect data"
		c.JSON(400, erMes)
	}
}
