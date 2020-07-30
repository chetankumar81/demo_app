package controllers

import (
	"demo_app/models"
	"demo_app/util"
	"encoding/json"
	"kbyp-common-libs/utility/log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/spf13/cast"
)

type StartGameRequest struct {
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}

type PickCardRequest struct {
	GameId     string `json:"gameId"`
	User       string `json:"user"`
	Card       string `json:"card"`
	PickedTime string `json:"pickedTime"`
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

//PickCard ...
func PickCard(request events.APIGatewayProxyRequest) (response string) {
	responseJson := util.ResponseJSON{}
	responseJson.Code = 400
	responseJson.Model = "Error in picking card"

	var pickCardRequest PickCardRequest
	err := json.Unmarshal([]byte(request.Body), &pickCardRequest)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

	user, err := models.GetUserByuserName(pickCardRequest.User)
	if err != nil {
		log.Println("Error in GetUserByuserName", err)
		return
	}
	game, err := models.GetGameById(cast.ToInt(pickCardRequest.GameId))
	if err != nil {
		log.Println("Error in GetGameById", err)
		return
	}
	card, err := models.GetCardMapById(cast.ToInt(pickCardRequest.Card))
	if err != nil {
		log.Println("Error in GetCardMapById", err)
		return
	}

	last3Cards, err := models.GetLast3CardValue(game.Id, user.Id)
	if err != nil {
		log.Println("Error in GetLast3Cards", err)
		return
	}

	cards := &models.Cards{}
	cards.GameId = game
	cards.UserId = user
	cards.Card = card
	cards.PickedTime = cast.ToTime(pickCardRequest.PickedTime)

	_, err = models.AddCards(cards)

	flag := false
	log.Println(last3Cards, card.CardVal)
	if card.CardVal > last3Cards[2] && last3Cards[0] > last3Cards[1] && last3Cards[1] > last3Cards[2] {
		flag = true
	}

	responseModel := make(map[string]interface{})
	if flag {
		game.Status = 0
		game.Ended = time.Now()

		err = models.UpdateGameById(game)
		responseModel["winner"] = user.UserName

	}

	responseJson.Msg = "Success"
	responseJson.Model = responseModel
	responseJson.Code = 200
	response = util.GetResponseJSONInString(responseJson)
	return

}
