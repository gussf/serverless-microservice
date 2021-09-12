package lambdas

import (
	"net/http"

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
		return FormatErrorMessageToAPIGatewayResponse(err, http.StatusOK), nil
	}

	if err != nil {
		return FormatErrorMessageToAPIGatewayResponse(err, http.StatusInternalServerError), nil
	}

	return FormatAPIGatewayResponse(cars, http.StatusOK), nil
}
