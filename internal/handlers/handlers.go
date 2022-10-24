package handlers

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	some := c.Param("id")
	c.String(200, some)
}
