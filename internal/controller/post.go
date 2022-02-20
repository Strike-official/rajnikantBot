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
	email_id := ctx.Query("email_id")
	user_name := ctx.Query("user_name")
	bot_name := ctx.Query("bot_name")
	strikeObj := core.CreateBot_3(request, bot_id, pic_url, email_id, user_name, bot_name)
	ctx.JSON(200, strikeObj)
}

func EditBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.EditBot(request)
	ctx.JSON(200, strikeObj)
}
func AddHandlerToBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.AddHandlerToBot(request)
	ctx.JSON(200, strikeObj)
}
