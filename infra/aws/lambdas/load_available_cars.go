package lambdas

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
	"github.com/gussf/serverless-microservice/domain/usecases"
	"github.com/gussf/serverless-microservice/infra/db"
)

func LoadAvailableCarsHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	stubCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, false)
	repo := db.StubRepository{CarList: []i.CarDAO{stubCar}}
	svc := usecases.NewLoadAvailableCarsService(repo)
	cars, err := svc.List()

	if err == e.ErrNoCarsAvailable {
		// return events.APIGatewayProxyResponse{
		// 	Body:       fmt.Sprintf("%v", ErrorMessage{Message: err.Error()}),
		// 	StatusCode: 200,
		// }, nil
		return events.APIGatewayProxyResponse{StatusCode: 200}, err
	}

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", cars),
		StatusCode: 200,
	}, nil
}
