package core

import (
	"fmt"

	"github.com/Strike-official/rajnikantBot/internal/model"
	pkg "github.com/Strike-official/rajnikantBot/pkg/mongodb"
	"github.com/strike-official/go-sdk/strike"
)

func YourBots(request model.Request_Structure) *strike.Response_structure {
	botNames := getAllBotsByUser(request)

	strikeObject := strike.Create("your_bot", model.Conf.APIEp+"/your_bot_1?userName=")

	question_object := strikeObject.Question("userName").QuestionText().
		SetTextToQuestion("Select which bot to edit?", "Text Description, getting used for testing purpose.")

	answer_object := question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)

	for _, botName := range botNames {
		answer_object = answer_object.AnswerCard().SetHeaderToAnswer(2, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H3, botName, "#e456tc", true)
	}

	return strikeObject
}

func YourBots_1(request model.Request_Structure) *strike.Response_structure {
	var r *strike.Response_structure

	return r
}

func getAllBotsByUser(request model.Request_Structure) []string {
	botNames := make(map[string]string)
	documentArray := pkg.GetByField("businessAccounts", "userid", request.Bybrisk_session_variables.UserId)
	for _, doc := range documentArray {

		botNames[fmt.Sprintf("%v", doc["username"]] = fmt.Sprintf("%v", doc["_id"])
	}
	return botNames
}
