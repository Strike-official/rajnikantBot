package controller

import (
	"fmt"

	"github.com/Strike-official/rajnikantBot/internal/core"
	"github.com/Strike-official/rajnikantBot/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.CreateBot(request)
	ctx.JSON(200, strikeObj)
}

func CreateBot_1(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.CreateBot_1(request)
	ctx.JSON(200, strikeObj)
}

func CreateBot_2(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	username := ctx.Query("userName")
	strikeObj := core.CreateBot_2(request, username)
	ctx.JSON(200, strikeObj)
}

func CreateBot_3(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	bot_id := ctx.Query("bot_id")
	pic_url := ctx.Query("pic_url")
	strikeObj := core.CreateBot_3(request, bot_id, pic_url)
	ctx.JSON(200, strikeObj)
}

func YourBots(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.YourBots(request)
	ctx.JSON(200, strikeObj)
}

func YourBots_1(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.YourBots_1(request, ctx.Query("botName"), ctx.Query("actionName"))
	ctx.JSON(200, strikeObj)
}

func YourBots_2(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.YourBots_2(request, ctx.Query("botName"), ctx.Query("actionName"))
	ctx.JSON(200, strikeObj)
}
