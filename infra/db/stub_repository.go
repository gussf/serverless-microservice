package db

import (
	"github.com/gussf/serverless-microservice/domain/errors"
	i "github.com/gussf/serverless-microservice/domain/interfaces"
)

type StubRepository struct {
	CarList []i.CarDAO
}

func (r StubRepository) ListCars() ([]i.CarDAO, error) {
	return r.CarList, nil
}

type StubRepositoryInternalError struct {
}

func (s StubRepositoryInternalError) ListCars() ([]i.CarDAO, error) {
	return nil, errors.ErrFetchingAvailableCars
}
