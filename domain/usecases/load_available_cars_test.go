package usecases

import (
	"testing"

	"github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
	"github.com/gussf/serverless-microservice/infra/db"
)

func TestLoadAvailableCars_NoCarsAvailable(t *testing.T) {

	unavailableCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, false)
	carList := []i.CarDAO{unavailableCar}
	r := db.StubRepository{CarList: carList}
	a := NewLoadAvailableCarsService(r)

	cars, err := a.List()

	if cars != nil {
		t.Errorf("cars - expected: nil, got: %v", cars)
	}

	if err != errors.ErrNoCarsAvailable {
		t.Errorf("err - expected: ErrNoCarsAvailable, got: %v", err)
	}
}
