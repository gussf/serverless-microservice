package db

import i "github.com/gussf/serverless-microservice/domain/interfaces"

type StubRepository struct {
	CarList []i.CarDAO
}

func (r StubRepository) ListCars() ([]i.CarDAO, error) {
	return r.CarList, nil
}
