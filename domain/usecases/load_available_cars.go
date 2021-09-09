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

	if len(cars) == 0 {
		return nil, errors.ErrNoCarsAvailable
	}

	return cars, nil
}