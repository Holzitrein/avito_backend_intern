package routers

import (
	"avito_balance/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/user/:id", handlers.GetBalance) //получение баланса
	//route.POST("/user/:id", handlers.AddBalance)      //добавление денег на баланс
	//route.POST("/user/:id", handlers.TransferBalance) //перевод между пользователями
	//route.POST("/user/:id", handlers.ReserveBalance)  //Резервирование денег
	//route.POST("/user/:id", handlers.ConformBalnce)   //Признание выручки
	//route.GET("/user/:id", handlers.ReportService)    //Отчёт по услуге за n период
	//route.GET("/user/:id", handlers.ReportOperation)  //Отчёт по операциям
}
