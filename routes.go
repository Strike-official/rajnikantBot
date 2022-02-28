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
		cb.POST("/your_bot", controller.YourBots)
		cb.POST("/your_bot_1", controller.YourBots_1)
		cb.POST("/your_bot_2", controller.YourBots_2)
	}
}
