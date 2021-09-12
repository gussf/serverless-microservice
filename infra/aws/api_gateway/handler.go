package api_gateway

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gussf/serverless-microservice/infra/aws/lambdas"
)

func Run() {
	lambda.Start(Handler)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Path == "/cars/available" && req.HTTPMethod == http.MethodGet {
		return lambdas.LoadAvailableCarsHandler(req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
	}, nil
}
