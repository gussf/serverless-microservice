package usecases

import (
	"github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
)

type LoadAvailableCarsService struct {
	repo i.Repository
}

func NewLoadAvailableCarsService(repo i.Repository) LoadAvailableCarsService {
	return LoadAvailableCarsService{repo}
}

func (s LoadAvailableCarsService) List() ([]i.CarDAO, error) {
	cars, err := s.repo.ListCars()

	if err != nil {
		return nil, errors.ErrFetchingAvailableCars
	}

	availableCars := GetAvailableCarsFrom(cars)

	if len(availableCars) == 0 {
		return nil, errors.ErrNoCarsAvailable
	}

	return availableCars, nil
}

func GetAvailableCarsFrom(cars []i.CarDAO) []i.CarDAO {
	var availableCars []i.CarDAO
	for _, c := range cars {
		if c.IsAvailable() {
			availableCars = append(availableCars, c)
		}
	}
	return availableCars

}
