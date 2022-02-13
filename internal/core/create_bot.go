package core

import (
	"fmt"
	"log"

	"github.com/Strike-official/rajnikantBot/internal/model"
	pkg "github.com/Strike-official/rajnikantBot/pkg/mongodb"

	sesMailer "github.com/Strike-official/rajnikantBot/pkg/sesMailer"
	tinyUrl "github.com/Strike-official/rajnikantBot/pkg/tinyurl"
	"github.com/strike-official/go-sdk/strike"
)

var (
	mailSubject = "Welcome👋 to the Strike Developer Community [🤖 %s]"
	htmlBody    = "<p>Hi,</p>" +
		"<p> Welcome to Strike developers community. And congratualtions on making your first application live on Strike platform.</p>" +
		"<p>Below are the details of your bot</p>" +
		"<ul>" +
		"<li>User Id: %s </li>" +
		"<li>Bot Name: %s </li>" +
		"<li>Bot ID: %s </li>" +
		"<li>Bot QR Code: Do share the QR or AppLink with your friends, and let them know how cool your app is 😎 :</li>" +
		"<img src=%s width='200' height='200'><br>" +
		"<li>Bot Link - <a href=%s>%s</a></li></ul>" +
		"<br>" +
		"<p>Currently your application is in development mode, please use rajnikant-bot to update the API Link of your bot. <br>Find the below resources on how to create an application which interfaces well with strike platform.</p>" +
		"<a href='https://bybrisk-strike.gitbook.io/strike-developer-community'>How to create app on Strike (Extensive documentation)</a><br>" +
		"<a href='https://bybrisk-strike.gitbook.io/strike-developer-community/go'>Create App with Strike-Go-SDK</a><br>" +
		"<a href='https://bybrisk-strike.gitbook.io/strike-developer-community/go-1'>Create App with Strike-Go-SDK</a><br>" +
		"<a href='https://bybrisk-strike.gitbook.io/strike-developer-community'>Update API of the application using Rajnikant</a><br>" +
		"<a href='https://github.com/Strike-official'>Github sample applications</a><br>" +
		"<p>Please reach out to, strike.info@bybrisk.com for any queries</p>" +
		"<br>" +
		"<p>Regards,<br>Strike</p>"
	textBody = "Welcome to Strike developers community. And congratualtions on making your first application live on Strike platform."
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
			SetTextToQuestion("Please provide the profile pic URL of the bot.", "Text Description, getting used for testing purpose.")

		question_object3.Answer(true).TextInput("Input Description")

		question_object4 := strikeObject.Question("emailId").QuestionText().
			SetTextToQuestion("What is your email? preferably the developer's email-id", "Text Description, getting used for testing purpose.")

		question_object4.Answer(true).TextInput("Input Description")

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
	strikeObject := strike.Create("getting_started", model.Conf.APIEp+"/create_3?bot_id="+id+"&pic_url="+request.User_session_variables.PicURL+"&email_id="+request.User_session_variables.EmailId)

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

func CreateBot_3(request model.Request_Structure, bot_id string, pic_url, email_id string) *strike.Response_structure {
	botLink := "https://bybrisk.page.link/?link=https://bybrisk.com?business_id=" + bot_id + "&apn=com.bybrisk.strike.app"
	botLinkOriginal := botLink
	l := tinyUrl.FetchTinyUrl(botLink)
	if l != "" {
		botLink = l
	}
	botQrLink := "https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=" + botLinkOriginal
	docLink := "https://bybrisk-strike.gitbook.io"
	strikeObject := strike.Create("getting_started", "")
	strikeObject.Question("").QuestionCard().
		SetHeaderToQuestion(1, strike.HALF_WIDTH).
		AddTextRowToQuestion(strike.H4, "Congractulations on your new bot. \n\n You can add your bot using this link "+botLink+" \n\nPlease see your bots section to edit your bot or create more Action Handlers. Read our docs here : "+email_id+docLink, "black", false).
		AddGraphicRowToQuestion(strike.PICTURE_ROW, []string{botQrLink}, []string{"tumbnail.jpeg"})
	addBotSchemaToMongo(request, bot_id, pic_url)
	mailSubject = fmt.Sprintf(mailSubject, request.User_session_variables.Title)
	htmlBody = fmt.Sprintf(htmlBody, request.Bybrisk_session_variables.UserId, request.User_session_variables.Title, bot_id, botQrLink, botLink, request.User_session_variables.Title)
	fmt.Println(htmlBody)
	sesMailer.SendMail(email_id, mailSubject, htmlBody, textBody)
	return strikeObject
}

func addBotAccountToMongo(request model.Request_Structure, userName string) string {

	botAccountInterface := &model.BotAccount{
		PicURL:           request.User_session_variables.PicURL,
		UserName:         userName,
		BusinessName:     request.User_session_variables.BusinessName,
		Email:            request.User_session_variables.EmailId,
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
