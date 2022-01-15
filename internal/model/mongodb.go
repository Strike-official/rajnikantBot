package model

type BotAccount struct {
	// The url of the profile pic for this Bot
	//
	PicURL string `json: "picurl"`

	// uniquely identify the bot
	//
	UserName string `json: "username"`

	// The email Id of the bot admin
	//
	Email string `json: "email"`
	// Display name of the bot
	//
	BusinessName string `json: "businessname"`
	// Category the bot belongs to
	//
	BusinessCategory string `json: "businessCat"`

	// Address of the of the bot admin
	//
	Address string `json: "address"`

	// Redundant field we need to remove in subsequent version. Currently we will just populate something arbitrarly
	//
	Latitude float64 `json:"latitude"`

	// Redundant field we need to remove in subsequent version. Currently we will just populate something arbitrarly
	//
	Longitude float64 `json:"longitude"`

	Bybid string `json:"bybid"`

	UserId string `json:"userid"`
}

type BotSchema struct {
	Businessid string           `json:"businessid"`
	Initialise InitialiseSchema `json:"initialise"`
}

type InitialiseSchema struct {
	QCard          QCardStructure            `json:"QCard"`
	ActionHandlers []ActionHandlersStructure `json:"actionHandlers"`
}

type QCardStructure struct {
	CardPic         string `json:"cardPic"`
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Story           string `json:"story"`
	TopRightData    string `json:"topRightData"`
	BottomRightData string `json:"bottomRightData"`
}

type ActionHandlersStructure struct {
	Handler      string             `json:"handler"`
	ObservedName string             `json:"observedName"`
	HandlerType  string             `json:"handlerType"`
	APIObject    APIObjectStructure `json:"apiObject"`
}

type APIObjectStructure struct {
	APIURL string `json:"apiURL"`
}
