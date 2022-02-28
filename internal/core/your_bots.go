package core

import (
	"fmt"
	"log"
	"strings"

	"github.com/Strike-official/rajnikantBot/internal/model"
	pkg "github.com/Strike-official/rajnikantBot/pkg/mongodb"
	"github.com/strike-official/go-sdk/strike"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func YourBots(request model.Request_Structure) *strike.Response_structure {
	botNames := getAllBotsByUser(request)

	strikeObject := strike.Create("your_bot", model.Conf.APIEp+"/your_bot_1?botName=NA&actionName=NA")

	question_object := strikeObject.Question("botUserName").QuestionText().
		SetTextToQuestion("Select which bot to edit?", "Text Description, getting used for testing purpose.")

	answer_object := question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION)

	for _, botName := range botNames {
		answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, botName, "#008f5a", false)
	}

	question_object2 := strikeObject.Question("actionOnBot").QuestionText().
		SetTextToQuestion("Which action to perform?", "Text Description, getting used for testing purpose.")

	question_object2.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Modify Bot Details", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Update Action Handlers", "#008f5a", false)
	//AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Delete Bot", "#008f5a", false)

	return strikeObject
}

func YourBots_1(request model.Request_Structure, botUsername, actionOnBot string) *strike.Response_structure {
	if botUsername != "NA" {
		return actionFuncMap[actionOnBot](botUsername)
	}
	return actionFuncMap[request.User_session_variables.ActionOnBot[0]](request)
}

func YourBots_2(request model.Request_Structure, botName string, actionName string) *strike.Response_structure {
	return actionFuncMap[actionName](botName, request)
}

// Helper functions

func modifyBotDetails(Values ...interface{}) *strike.Response_structure {
	var strikeObject *strike.Response_structure
	switch v := Values[0].(type) {
	case string:
		strikeObject = strike.Create("your_bot", model.Conf.APIEp+"/your_bot_2?botName="+v+"&actionName=updateBotDetailByField")
	case model.Request_Structure:
		strikeObject = strike.Create("your_bot", model.Conf.APIEp+"/your_bot_2?botName="+v.User_session_variables.BotUserName[0]+"&actionName=updateBotDetailByField")
	}

	question_object := strikeObject.Question("modifyDetailField").QuestionText().
		SetTextToQuestion("Select the field to modify", "Text Description, getting used for testing purpose.")

	question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Card Title", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Card Subtitle", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Card Story", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Card Cover Photo", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Bot Display Name", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Bot Category", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Bot Profile Pic", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Developer's Email", "#008f5a", false)

	question_object2 := strikeObject.Question("modifyDetailValue").QuestionText().
		SetTextToQuestion("Please paste the updated value", "Text Description, getting used for testing purpose.")
	question_object2.Answer(true).TextInput("Input Description")

	return strikeObject
}

func updateBotDetailByField(Value ...interface{}) *strike.Response_structure {
	var action string
	val := Value[1].(model.Request_Structure).User_session_variables
	switch action = val.ModifyDetailField[0]; action {
	case "Card Title":
		updateQCardDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "initialise.qcard.title")
	case "Card Subtitle":
		updateQCardDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "initialise.qcard.subtitle")
	case "Card Story":
		updateQCardDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "initialise.qcard.story")
	case "Card Cover Photo":
		updateQCardDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "initialise.qcard.cardpic")
	case "Bot Display Name":
		updateBotBusinessDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "businessname")
	case "Bot Category":
		updateBotBusinessDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "businesscategory")
	case "Bot Profile Pic":
		updateBotBusinessDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "picurl")
	case "Developer's Email":
		updateBotBusinessDetailByActionHelper(Value[0].(string), val.ModifyDetailValue, "email")
	}

	strikeObject := strike.Create("your_bot", model.Conf.APIEp+"/your_bot_1?botName="+Value[0].(string)+"&actionName=Modify Bot Details")
	question_object := strikeObject.Question("na").QuestionText().
		SetTextToQuestion(action+" Updated!", "Text Description, getting used for testing purpose.")
	question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "↩ Back", "#008f5a", false)
	return strikeObject
}

func updateActionHandlers(Values ...interface{}) *strike.Response_structure {
	var botUserName string
	switch v := Values[0].(type) {
	case string:
		botUserName = v
	case model.Request_Structure:
		botUserName = Values[0].(model.Request_Structure).User_session_variables.BotUserName[0]
	}

	strikeObject := strike.Create("your_bot", model.Conf.APIEp+"/your_bot_2?botName="+botUserName+"&actionName=updateActionHandlersHelper")
	question_object := strikeObject.Question("actionHandlerName").QuestionText().
		SetTextToQuestion("Which Action Handler to modify?", "Text Description, getting used for testing purpose.")

	answer_object := question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION)

	actionHanlders := getActionHanldersHelper(botUserName)
	for _, ac := range actionHanlders {
		answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, ac, "#008f5a", false)
	}

	question_object1 := strikeObject.Question("actionOnHandler").QuestionText().
		SetTextToQuestion("Select action to perform", "Text Description, getting used for testing purpose.")

	question_object1.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Modify Handler Name", "#008f5a", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Update API Endpoint", "#008f5a", false)

	question_object2 := strikeObject.Question("modifyDetailValue").QuestionText().
		SetTextToQuestion("What's the new value?", "Text Description, getting used for testing purpose.")

	question_object2.Answer(true).TextInput("Input Description")
	return strikeObject
}

func updateActionHandlersHelper(Values ...interface{}) *strike.Response_structure {
	var action string
	val := Values[1].(model.Request_Structure).User_session_variables
	switch action = val.ActionOnHandler[0]; action {
	case "Update API Endpoint":
		updateActionHandlerArrayHelper(Values[0].(string), val.ModifyDetailValue, val.ActionHandlerName[0], "initialise.actionhandlers.$.apiobject.apiurl")
	case "Modify Handler Name":
		modifyActionHandlerName(Values[0].(string), val.ModifyDetailValue, val.ActionHandlerName[0], "initialise.actionhandlers.$.observedname")
	}

	nPrefix := strings.ReplaceAll(strings.ReplaceAll(action, "Modify", ""), "Update", "")

	strikeObject := strike.Create("your_bot", model.Conf.APIEp+"/your_bot_1?botName="+Values[0].(string)+"&actionName=Update Action Handlers")
	question_object := strikeObject.Question("na").QuestionText().
		SetTextToQuestion(nPrefix+" Updated!", "Text Description, getting used for testing purpose.")
	question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "↩ Back", "#008f5a", false)
	return strikeObject
}

func deleteBot(Values ...interface{}) *strike.Response_structure {
	//Delete the bot from our system
	log.Println("Deleting bot : ", Values[0])
	var r *strike.Response_structure
	return r

}

var actionFuncMap map[string]func(...interface{}) *strike.Response_structure = map[string]func(...interface{}) *strike.Response_structure{
	"Update Action Handlers":     updateActionHandlers,
	"Modify Bot Details":         modifyBotDetails,
	"Delete bot":                 deleteBot,
	"updateBotDetailByField":     updateBotDetailByField,
	"updateActionHandlersHelper": updateActionHandlersHelper,
}

// DB helper functions
func getAllBotsByUser(request model.Request_Structure) []string {
	var botNames []string
	documentArray := pkg.GetByField("businessAccounts", "userid", request.Bybrisk_session_variables.UserId)
	for _, doc := range documentArray {
		botNames = append(botNames, fmt.Sprintf("%v", doc["username"]))
	}
	return botNames
}

func getActionHanldersHelper(botUserName string) []string {
	var actionHandlers []string
	botDocumentArray := pkg.GetByField("bot-schema", "businessid", getBotBusinessIdByBotname(botUserName))
	actionHandlersArray := botDocumentArray[0]["initialise"].(primitive.M)["actionhandlers"]

	for _, doc := range actionHandlersArray.(primitive.A) {
		actionHandlers = append(actionHandlers, doc.(primitive.M)["observedname"].(string))
	}
	return actionHandlers
}

func getBotBusinessIdByBotname(botUserName string) string {
	accountDocumentArray := pkg.GetByField("businessAccounts", "username", botUserName)
	return accountDocumentArray[0]["_id"].(primitive.ObjectID).Hex()
}

func updateQCardDetailByActionHelper(botUsername, modifiedValue, modifiedField string) {
	filter := bson.M{"businessid": getBotBusinessIdByBotname(botUsername)}
	update := bson.M{"$set": bson.M{modifiedField: modifiedValue}}
	if _, err := pkg.Update("bot-schema", filter, update); err != nil {
		log.Println("Error updating bot details: ", err)
	}
}

func updateBotBusinessDetailByActionHelper(botUsername, modifiedValue, modifiedField string) {
	filter := bson.M{"username": botUsername}
	update := bson.M{"$set": bson.M{modifiedField: modifiedValue}}
	if _, err := pkg.Update("businessAccounts", filter, update); err != nil {
		log.Println("Error updating bot details: ", err)
	}
}

func updateActionHandlerArrayHelper(botUsername, modifiedValue, observedName, modifiedField string) {
	filter := bson.M{"businessid": getBotBusinessIdByBotname(botUsername), "initialise.actionhandlers.observedname": observedName}
	update := bson.M{"$set": bson.M{modifiedField: modifiedValue}}
	if _, err := pkg.Update("bot-schema", filter, update); err != nil {
		log.Println("Error updating Action Handlers: ", err)
	}
}

func modifyActionHandlerName(botUsername, modifiedValue, observedName, modifiedField string) {
	filter := bson.M{"businessid": getBotBusinessIdByBotname(botUsername), "initialise.actionhandlers.observedname": observedName}
	update := bson.M{"$set": bson.M{modifiedField: modifiedValue}}
	if _, err := pkg.Update("bot-schema", filter, update); err != nil {
		log.Println("Error updating Action Handlers: ", err)
	}
}
