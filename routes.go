package main

import (
	"github.com/Strike-official/rajnikantBot/internal/controller"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	cb := router.Group("/rajnikantBot")
	{
		cb.POST("/create", controller.CreateBot)
		cb.POST("/create_1", controller.CreateBot_1)
		cb.POST("/create_2", controller.CreateBot_2)
		cb.POST("/create_3", controller.CreateBot_3)
		cb.POST("/edit", controller.EditBot)
		cb.POST("/add", controller.AddHandlerToBot)
	}
}
