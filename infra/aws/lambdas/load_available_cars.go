package lambdas

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
	"github.com/gussf/serverless-microservice/domain/usecases"
	"github.com/gussf/serverless-microservice/infra/db"
)

func LoadAvailableCarsHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// temporary
	stubCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, true)
	stubCar_2 := i.NewCarDAO("Uno", "Fiat", 8000000, true)

	repo := db.StubRepository{CarList: []i.CarDAO{stubCar, stubCar_2}}
	svc := usecases.NewLoadAvailableCarsService(repo)
	cars, err := svc.List()

	if err == e.ErrNoCarsAvailable {
		msg, _ := json.Marshal(ErrorMessage{Message: err.Error()})
		return events.APIGatewayProxyResponse{
			Body:       string(msg),
			StatusCode: 200,
		}, nil
	}

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	msg, err := json.Marshal(cars)
	return events.APIGatewayProxyResponse{
		Body:       string(msg),
		StatusCode: 200,
	}, nil
}
