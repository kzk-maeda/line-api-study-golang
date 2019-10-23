package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	line := Line{}
	err := line.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("ChannelSecret : " + line.ChannelSecret)
	// fmt.Println("ChannelToken : " + line.ChannelToken)
	// fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	// fmt.Printf("Body = %d.\n", request.Body)
	// fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}
	eve, err := ParseRequest(line.ChannelSecret, request)
	if err != nil {
		status := 200
		if err == linebot.ErrInvalidSignature {
			status = 400
		} else {
			status = 500
		}
		return events.APIGatewayProxyResponse{StatusCode: status}, errors.New("Bad Request")
	}
	line.EventRouter(eve)
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
