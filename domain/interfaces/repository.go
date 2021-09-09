package interfaces

type CarDAO struct {
	name      string
	brand     string
	price     float64
	available bool
}

func NewCarDAO(name string, brand string, price float64, available bool) CarDAO {
	return CarDAO{
		name, brand, price, available,
	}
}

func (c CarDAO) Available() bool {
	return c.available
}

type Repository interface {
	ListCars() ([]CarDAO, error)
}
