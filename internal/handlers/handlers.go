package handlers

import (
	"avito_balance/internal/database"
	"avito_balance/internal/model"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	var requestBody model.BalanceGet
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	reply := database.GetBalanceDb(requestBody)
	c.JSON(200, reply)
}

func TransferBalance(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func AddBalance(c *gin.Context) {
	var requestBody model.BalanceAdd
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	reply := database.AddBalanceDb(requestBody)
	c.JSON(200, reply)
}

func ReserveBalance(c *gin.Context) {
	var requestBody model.ReserveCreate
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "error")
	}
	reply := database.ReserveBalanceDb(requestBody)
	c.JSON(200, reply)
}

func ConformBalnce(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func ReportService(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func ReportOperation(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}
