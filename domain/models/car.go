package models

type Car struct {
	name        string
	brand       string
	price_cents float64
	available   bool
}

func NewCar(name string, brand string, price_cents float64, available bool) Car {
	return Car{
		name, brand, price_cents, available,
	}
}

func (c Car) Available() bool {
	return c.available
}
