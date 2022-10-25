package routers

import (
	"avito_balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/user", handlers.GetBalance)                //получение баланса
	route.POST("/user", handlers.AddBalance)               //добавление денег на баланс
	route.POST("/user/transfer", handlers.TransferBalance) //перевод между пользователями
	route.POST("/order/create", handlers.ReserveBalance)   //Резервирование денег
	route.POST("/order/", handlers.ConformBalnce)          //Признание выручки или отклонение заказа
	route.GET("/service", handlers.ReportService)          //Отчёт по услуге за n период
	//route.GET("/user/:id", handlers.ReportOperation)  //Отчёт по операциям
}
