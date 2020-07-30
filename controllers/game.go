package controllers

import (
	"demo_app/models"
	"demo_app/util"
	"encoding/json"
	"kbyp-common-libs/utility/log"

	"github.com/aws/aws-lambda-go/events"
)

type StartGameRequest struct {
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}

//StartGame ...
func StartGame(request events.APIGatewayProxyRequest) (response string) {
	responseJson := util.ResponseJSON{}
	responseJson.Code = 400
	responseJson.Model = "Error in StartGame"

	var startGameRequest StartGameRequest
	err := json.Unmarshal([]byte(request.Body), &startGameRequest)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

	user1, err := models.GetUserByuserName(startGameRequest.User1)
	if err != nil {
		log.Println("Error in GetUserByuserName", err)
		return
	}
	user2, err := models.GetUserByuserName(startGameRequest.User2)
	if err != nil {
		log.Println("Error in GetUserByuserName", err)
		return
	}

	if user1 == nil || user2 == nil {
		responseJson.Model = "Invalid userIds"
		return
	}

	game := &models.Game{}
	game.User1 = user1
	game.User2 = user2
	game.Status = 1
	game.Timer = "30"

	gameId, err := models.AddGame(game)
	if err != nil {
		log.Println("Error in AddGame", err)
		return
	}

	responseModel := make(map[string]interface{})
	responseModel["gameId"] = gameId

	responseJson.Model = responseModel
	responseJson.Code = 200
	response = util.GetResponseJSONInString(responseJson)
	return
}
