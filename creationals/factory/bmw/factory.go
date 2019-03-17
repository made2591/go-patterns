package bmw

import (
	"errors"
	"fmt"

	"github.com/made2591/go-patterns/creationals/factory/factory"
	"github.com/made2591/go-patterns/creationals/factory/vehicle"
)

type bmwFactory struct {
}

func NewBmwFactory() factory.Factory {
	return &bmwFactory{}
}

func (f *bmwFactory) CreateVehicle(tipe uint) (vehicle.Vehicle, error) {
	switch tipe {
	case 1:
		return createBMWSerieOne(), nil
	case 2:
		return createBMWSerieTwo(), nil
	default:
		return nil, errors.New(fmt.Sprintf("BMW tipe %d not supported", tipe))
	}
}
