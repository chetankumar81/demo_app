package main

import (
	"context"
	"demo_app/db"
	"demo_app/handlers"
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/spf13/viper"
)

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(lambdactx context.Context, req map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
	}

	awsResponse := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		Body:            `{"code":"512","msg":"` + viper.GetString("maintenance.512") + `","model":{}}`,
		Headers:         headers,
		StatusCode:      200,
	}

	xray.Configure(xray.Config{
		LogLevel:       "info",
		ServiceVersion: "1.2.3",
	})

	if !db.IsDbConnected("gl") {
		log.Print("Database not connected")
		awsResponse.StatusCode = 500
		awsResponse.Body = `{ "message": "Server is busy, please try again later" }`
		return awsResponse, nil
	}

	var request events.APIGatewayProxyRequest

	//marshal the request to json
	requestBytes, _ := json.Marshal(req)
	var returnString string
	var err error
	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		log.Print("Error in unmarshaling the request", err)
		return awsResponse, nil
	}

	log.Print("Path:", request.Path, "HTTPMethod:", request.HTTPMethod)
	log.Print("HTTPRequest:", request.Body)

	switch true {
	case strings.Contains(request.Path, "/start_game"):
		returnString = handlers.StartGame(request)
		break

	case strings.Contains(request.Path, "/pick_card"):
		returnString = handlers.PickCard(request)
		break

	case strings.Contains(request.Path, "/game_details"):
		returnString = handlers.GetGameDetails(request)
		break
	default:
		log.Print("Path did not match")
		break
	}

	awsResponse.Body = returnString
	log.Print("response", awsResponse)
	log.Print("@EOR@")
	return awsResponse, nil
}
