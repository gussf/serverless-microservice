package lambdas

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/gussf/serverless-microservice/domain/errors"
	"github.com/gussf/serverless-microservice/domain/usecases"
	"github.com/gussf/serverless-microservice/infra/db"
)

func LoadAvailableCarsHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	repo2 := db.NewDynamoRepository()
	fmt.Println(repo2)
	// temporary
	// stubCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, true)
	// stubCar_2 := i.NewCarDAO("Uno", "Fiat", 8000000, true)
	// stubCarList := []i.CarDAO{stubCar, stubCar_2}

	// repo := db.StubRepository{CarList: stubCarList}

	svc := usecases.NewLoadAvailableCarsService(repo2)
	cars, err := svc.List()

	if err == e.ErrNoCarsAvailable {
		return FormatErrorMessageToAPIGatewayResponse(err, http.StatusOK), nil
	}

	if err != nil {
		return FormatErrorMessageToAPIGatewayResponse(err, http.StatusInternalServerError), nil
	}

	return FormatAPIGatewayResponse(cars, http.StatusOK), nil
}
