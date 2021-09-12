package models

type Car struct {
	name        string
	brand       string
	price_cents int64
	available   bool
}

func NewCar(name string, brand string, price_cents int64, available bool) Car {
	return Car{
		name, brand, price_cents, available,
	}
}

func (c Car) IsAvailable() bool {
	return c.available
}
