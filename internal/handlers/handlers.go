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

func GetBalance(c *gin.Context) { //Получение баланса
	var requestBody model.BalanceGet
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, "error")
	}
	reply := database.GetBalanceDb(requestBody)
	c.JSON(200, reply)
}

// @Summary     Transfer balance
// @Description send money from one user to another
// @Tags        operation for user balance
// @Accept      json
// @Produce     json
// @Param       json model.TransferBalanceUser true "json struct for request"
// @Success     200  {object}                  model.TransferBalanceUserReturn
// @Failure     400  {object}                  model.ErrorReq
// @Router      /user/transfer [post]

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

// @Summary     add balance
// @Description Add money to user's balance
// @Tags        operation for user balance
// @Accept      json
// @Produce     json
// @Param       json model.BalanceAdd true "json struct for request"
// @Success     200  {object}         model.BalanceAddReturn
// @Failure     400  {object}         model.ErrorReq
// @Router      /user [post]

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

// @Summary     Reserve money
// @Description Request for reserve money
// @Tags        Reserve operation
// @Accept      json
// @Produce     json
// @Param       json model.ReserveCreate true "json struct for request"
// @Success     200  {object}            model.ReserveCreateReturn
// @Failure     400  {object}            model.ErrorReq
// @Router      /order/create [post]

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

// @Summary     Confirm or deny reserve
// @Description Request for confirm or deny reserve in order and balance
// @Tags        Reserve operation
// @Accept      json
// @Produce     json
// @Param       json model.ReserveConfirm true "json struct for request"
// @Success     200  {object}             model.ReserveCreateReturn
// @Failure     400  {object}             model.ErrorReq
// @Router      /order [post]

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

// @Summary     Get url for file
// @Description generate csv file and return url for file
// @Tags        Reports
// @Accept      json
// @Produce     json
// @Param       json model.ReportServiceStruct true "json struct for request"
// @Success     200  {object}                  model.ReportServiceStructReturn
// @Failure     400  {object}                  model.ErrorReq
// @Router      /service [get]

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

// @Summary     Get file
// @Description take url of file and return file for download
// @Tags        Reports
// @Accept      json
// @Produce     plain
// @Param       json model.ReportServiceStruct true   "url of file for request"
// @Success     200  {string}                  string "CSV file"
// @Failure     400  {object}                  model.ErrorReq
// @Router      /csv/{file} [get]

func ReportServiceFile(c *gin.Context) {
	url := c.Param("file")
	if _, err := os.Stat("internal/database/csv/" + url); os.IsNotExist(err) {
		c.JSON(404, "File not gound")
	}
	c.File("internal/database/csv/" + url)
}

// @Summary     Get report for operation
// @Description return report with array history of user
// @Tags        Reports
// @Accept      json
// @Produce     json
// @Param       json model.ReportOperationRequest true "json struct for request"
// @Success     200  {object}                     model.ReportOperationRequestTemp
// @Failure     400  {object}                     model.ErrorReq
// @Router      /user/history [get]

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
