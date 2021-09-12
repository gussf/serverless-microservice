package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	l "github.com/gussf/serverless-microservice/infra/aws/lambdas"
)

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Path == "/cars/available" {
		return l.LoadAvailableCarsHandler(req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
	}, nil
}

func main() {
	fmt.Println("MAIN")
	lambda.Start(handlers)
}
