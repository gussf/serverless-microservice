package usecases

import (
	"testing"

	"github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
)

type StubRepository struct {
	carList []i.CarDAO
}

func (r StubRepository) ListCars() ([]i.CarDAO, error) {
	var unavailableCars []i.CarDAO
	for _, c := range r.carList {
		if c.Available() {
			unavailableCars = append(unavailableCars, c)
		}
	}
	return unavailableCars, nil
}

func TestLoadAvailableCars_NoCarsAvailable(t *testing.T) {

	unavailableCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, false)
	carList := []i.CarDAO{unavailableCar}
	r := StubRepository{carList: carList}
	a := NewLoadAvailableCarsService(r)

	cars, err := a.List()

	if cars != nil {
		t.Errorf("cars - expected: nil, got: %v", cars)
	}

	if err != errors.ErrNoCarsAvailable {
		t.Errorf("err - expected: ErrNoCarsAvailable, got: %v", err)
	}
}
