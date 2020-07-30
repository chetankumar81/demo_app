package handlers

import (
	"demo_app/controllers"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

//StartGame ... func
func StartGame(request events.APIGatewayProxyRequest) (response string) {
	log.Print("starting game")
	response = "Error in starting Game"

	//checking the request method
	switch true {
	case request.HTTPMethod == "POST":
		response = controllers.StartGame(request)
	}
	return
}

//PickCard ... func
func PickCard(request events.APIGatewayProxyRequest) (response string) {
	log.Print("Card pick API start")
	response = "Error in Picking card"

	switch true {
	case request.HTTPMethod == "POST":
		response = controllers.PickCard(request)
	}
	return
}

//GetGameDetails ...
func GetGameDetails(request events.APIGatewayProxyRequest) (response string) {
	return "Getting Game Details"
}
