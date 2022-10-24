package routers

import (
	"avito_balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/user/:id", handlers.Hello)  //получение баланса
	route.POST("/user/:id", handlers.Hello) //добавление денег на баланс
	route.GET("/user/:id", handlers.Hello)  //Резервирование денег
	route.GET("/user/:id", handlers.Hello)  //Признание выручки
	route.GET("/user/:id", handlers.Hello)  //Отчёт по услуге за n период
	route.GET("/user/:id", handlers.Hello)  //Отчёт по операциям
}
