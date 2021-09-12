package lambdas

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func FormatAPIGatewayResponse(v interface{}, httpCode int) events.APIGatewayProxyResponse {
	jsonMsg, _ := json.Marshal(v)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonMsg),
		StatusCode: httpCode,
	}
}

func FormatErrorMessageToAPIGatewayResponse(err error, httpCode int) events.APIGatewayProxyResponse {
	errMsg := ErrorMessage{Message: err.Error()}
	jsonMsg, _ := json.Marshal(errMsg)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonMsg),
		StatusCode: httpCode,
	}
}
