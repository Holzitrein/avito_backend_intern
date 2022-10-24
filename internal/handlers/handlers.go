package handlers

import "github.com/gin-gonic/gin"

func GetBalance(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func TransferBalance(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func AddBalance(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}

func ReserveBalance(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
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
