package factory

import (
	"github.com/made2591/go-patterns/creationals/factory/vehicle"
)

type Factory interface {
	CreateVehicle(tipe uint) (vehicle.Vehicle, error)
}
