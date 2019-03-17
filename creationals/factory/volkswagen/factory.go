package volkswagen

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/factory/factory"
	"github.com/made2591/go-patterns/creationals/factory/vehicle"
)

type volkswagenFactory struct {
}

func NewVolkswagenFactory() factory.Factory {
	return &volkswagenFactory{}
}

func (f *volkswagenFactory) CreateVehicle(tipe uint) (vehicle.Vehicle, error) {
	switch tipe {
	case 1:
		return createVolkswagenPolo(), nil
	case 2:
		return createVolkswagenGolf(), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Volkswagen tipe %d not supported", tipe))
	}
}
