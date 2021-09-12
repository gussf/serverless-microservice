package interfaces

type CarDAO struct {
	Name      string
	Brand     string
	Price     int64
	Available bool
}

func NewCarDAO(name string, brand string, price int64, available bool) CarDAO {
	return CarDAO{
		name, brand, price, available,
	}
}

func (c CarDAO) IsAvailable() bool {
	return c.Available
}

type Repository interface {
	ListCars() ([]CarDAO, error)
}
