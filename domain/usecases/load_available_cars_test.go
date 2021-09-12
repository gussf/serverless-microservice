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
	svc := NewLoadAvailableCarsService(r)

	cars, err := svc.List()

	if cars != nil {
		t.Errorf("[]cars - expected: <nil>, got: %v", cars)
	}

	if err != errors.ErrNoCarsAvailable {
		t.Errorf("err - expected: ErrNoCarsAvailable, got: %v", err)
	}
}

func TestLoadAvailableCars_Success(t *testing.T) {

	availableCar := i.NewCarDAO("Corsa", "Chevrolet", 1500000, true)
	availableCar_2 := i.NewCarDAO("Fiat", "Uno", 200000, true)
	carList := []i.CarDAO{availableCar, availableCar_2}
	expected := len(carList)

	r := db.StubRepository{CarList: carList}
	svc := NewLoadAvailableCarsService(r)

	actual_cars, err := svc.List()

	if err != nil {
		t.Errorf("err - expected: <nil>, got: %v", err)
	}

	if len(actual_cars) != expected {
		t.Errorf("[]cars - expected len: %d, got: %v", expected, actual_cars)
	}

	if actual_cars[0].Name != availableCar.Name {
		t.Errorf("[]cars - expected car name: %s, got: %v", availableCar.Name, actual_cars[0].Name)
	}

	if actual_cars[0].Brand != availableCar.Brand {
		t.Errorf("[]cars - expected car brand: %s, got: %v", availableCar.Brand, actual_cars[0].Brand)
	}

	if actual_cars[0].Price != availableCar.Price {
		t.Errorf("[]cars - expected car price: %d, got: %v", availableCar.Price, actual_cars[0].Price)
	}

	if actual_cars[0].IsAvailable() != availableCar.IsAvailable() {
		t.Errorf("[]cars - expected availabilty: %v, got: %v", availableCar.IsAvailable(), actual_cars[0].IsAvailable())
	}
}

func TestLoadAvailableCars_InternalError(t *testing.T) {

	r := db.StubRepositoryInternalError{}
	svc := NewLoadAvailableCarsService(r)

	actual_cars, err := svc.List()

	if actual_cars != nil {
		t.Errorf("[]cars - expected: <nil>, got: %v", actual_cars)
	}

	if err != errors.ErrFetchingAvailableCars {
		t.Errorf("err - expected: ErrFetchingAvailableCars, got: %v", err)
	}
}
