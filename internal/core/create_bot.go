package core

import (
	"log"

	"github.com/Strike-official/rajnikantBot/internal/model"
	pkg "github.com/Strike-official/rajnikantBot/pkg/mongodb"
	"github.com/strike-official/go-sdk/strike"
)

func CreateBot(request model.Request_Structure) *strike.Response_structure {

	strikeObject := strike.Create("getting_started", model.Conf.APIEp+"/create_1")

	question_object1 := strikeObject.Question("userName").QuestionText().
		SetTextToQuestion("Alright "+request.Bybrisk_session_variables.Username+", a new bot. Please choose a unique username for your bot.", "Text Description, getting used for testing purpose.")

	question_object1.Answer(true).TextInput("Input Description")

	return strikeObject
}

func CreateBot_1(request model.Request_Structure) *strike.Response_structure {
	var strikeObject *strike.Response_structure
	if isUserUnique(request.User_session_variables.UserName) {

		//Add data to mongo

		strikeObject = strike.Create("getting_started", model.Conf.APIEp+"/create_2?userName="+request.User_session_variables.UserName)

		question_object1 := strikeObject.Question("businessName").QuestionText().
			SetTextToQuestion("Good! How are we going to call it? Please choose a name.", "Text Description, getting used for testing purpose.")

		question_object1.Answer(true).TextInput("Input Description")

		question_object2 := strikeObject.Question("businessCategory").QuestionText().
			SetTextToQuestion("Great! What category your bot belongs? Please name an industry.", "Text Description, getting used for testing purpose.")

		question_object2.Answer(true).TextInput("Input Description")

		question_object3 := strikeObject.Question("picURL").QuestionText().
			SetTextToQuestion("One last thing! Please provide the profile pic URL of the bot.", "Text Description, getting used for testing purpose.")

		question_object3.Answer(true).TextInput("Input Description")
		return strikeObject
	} else {
		strikeObject = strike.Create("getting_started", model.Conf.APIEp+"/create_1")

		question_object1 := strikeObject.Question("userName").QuestionText().
			SetTextToQuestion("Sorry, this username is already taken. Please try something different.", "Text Description, getting used for testing purpose.")

		question_object1.Answer(true).TextInput("Input Description")
		return strikeObject
	}

	return strikeObject
}

func CreateBot_2(request model.Request_Structure, userName string) *strike.Response_structure {

	// Create bot account
	id := addBotAccountToMongo(request, userName)

	// Create bot schema
	strikeObject := strike.Create("getting_started", model.Conf.APIEp+"/create_3?bot_id="+id+"&pic_url="+request.User_session_variables.PicURL)

	question_object1 := strikeObject.Question("title").QuestionCard().
		SetHeaderToQuestion(1, strike.HALF_WIDTH).
		AddGraphicRowToQuestion(strike.PICTURE_ROW, []string{"https://raw.githubusercontent.com/Strike-official/rajnikantBot/main/connecting.jpg"}, []string{"tumbnail.jpeg"}).
		AddTextRowToQuestion(strike.H4, "It's not 𝐜𝐨𝐧𝐧𝐞𝐜𝐭𝐢𝐧𝐠 𝐩𝐞𝐨𝐩𝐥𝐞 logo, it's 𝐩𝐥𝐞𝐚𝐬𝐞 𝐝𝐨𝐧'𝐭 𝐠𝐨 𝐈 𝐬𝐭𝐢𝐥𝐥 𝐠𝐨𝐭 𝐬𝐨𝐦𝐞 𝐪𝐮𝐞𝐬𝐭𝐢𝐨𝐧𝐬 𝐥𝐞𝐟𝐭 sign 😅", "black", false).
		AddTextRowToQuestion(strike.H4, "Now let's personalize your bot's greeting card", "black", false).
		AddTextRowToQuestion(strike.H4, " ", "black", false).
		AddGraphicRowToQuestion(strike.PICTURE_ROW, []string{"https://raw.githubusercontent.com/Strike-official/rajnikantBot/main/dazzle.png"}, []string{"tumbnail.jpeg"}).
		AddTextRowToQuestion(strike.H4, "What should be the title?", "black", false)
	question_object1.Answer(true).TextInput("Input Description")

	question_object2 := strikeObject.Question("subtitle").QuestionText().
		SetTextToQuestion("What should be the subtitle of the greeting card?", "Text Description, getting used for testing purpose.")

	question_object2.Answer(true).TextInput("Input Description")

	question_object3 := strikeObject.Question("story").QuestionText().
		SetTextToQuestion("What should be the story of the greeting card?", "Text Description, getting used for testing purpose.")

	question_object3.Answer(true).TextInput("Input Description")

	return strikeObject
}

func CreateBot_3(request model.Request_Structure, bot_id string, pic_url string) *strike.Response_structure {
	botLink := "https://bybrisk.page.link/?link=https://bybrisk.com?business_id=" + bot_id + "&apn=com.bybrisk.strike.app"
	docLink := "https://bybrisk-strike.gitbook.io"
	strikeObject := strike.Create("getting_started", "")
	strikeObject.Question("").QuestionCard().
		SetHeaderToQuestion(1, strike.HALF_WIDTH).
		AddTextRowToQuestion(strike.H4, "Congractulations on your new bot. You can add your bot using this link : "+botLink+" Please see your bots section to edit your bot or create more Action Handlers. Read our docs here : "+docLink, "black", false)
	addBotSchemaToMongo(request, bot_id, pic_url)
	return strikeObject
}

func addBotAccountToMongo(request model.Request_Structure, userName string) string {

	botAccountInterface := &model.BotAccount{
		PicURL:           request.User_session_variables.PicURL,
		UserName:         userName,
		BusinessName:     request.User_session_variables.BusinessName,
		Email:            "Not Provided",
		BusinessCategory: request.User_session_variables.BusinessCategory,
		Address:          "NA",
		Latitude:         23.2310402,
		Longitude:        77.45816359999999,
		UserId:           request.Bybrisk_session_variables.UserId,
	}

	botId := pkg.Insert("businessAccounts", botAccountInterface)
	return botId
}

func addBotSchemaToMongo(request model.Request_Structure, businessID string, cardPic string) {

	botSchemaInterface := &model.BotSchema{
		Businessid: businessID,
		Initialise: model.InitialiseSchema{
			QCard: model.QCardStructure{
				CardPic:         cardPic,
				Title:           request.User_session_variables.Title,
				Subtitle:        request.User_session_variables.Subtitle,
				Story:           request.User_session_variables.Story,
				TopRightData:    " ",
				BottomRightData: " ",
			},
			ActionHandlers: []model.ActionHandlersStructure{
				{
					Handler:      "getting-started",
					ObservedName: "Getting Started",
					HandlerType:  "API",
					APIObject: model.APIObjectStructure{
						APIURL: "http://ec2-18-218-96-97.us-east-2.compute.amazonaws.com/global-getting-started/getting-started",
					},
				},
			},
		},
	}

	id := pkg.Insert("bot-schema", botSchemaInterface)
	log.Println(id)
}

func isUserUnique(userName string) bool {
	return true
}
