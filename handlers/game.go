package handlers

import "github.com/aws/aws-lambda-go/events"

//StartGame ... func
func StartGame(request events.APIGatewayProxyRequest) (response string) {
	return "starting game"
}

//PickCard ... func
func PickCard(request events.APIGatewayProxyRequest) (response string) {
	return "Picking card"
}

//GetGameDetails ...
func GetGameDetails(request events.APIGatewayProxyRequest) (response string) {
	return "Getting Game Details"
}
