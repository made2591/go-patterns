package volkswagen

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/factory/vehicle"
)

const (
	PoloModel = "Polo"
	GolfModel = "Golf"
)

type volkswagenVehicle struct {
	model interface{}
}

func createVolkswagenPolo() vehicle.Vehicle {
	return &volkswagenVehicle{
		model: PoloModel,
	}
}

func createVolkswagenGolf() vehicle.Vehicle {
	return &volkswagenVehicle{
		model: GolfModel,
	}
}

func (v *volkswagenVehicle) GetModel() interface{} {
	return v.model
}

func (v *volkswagenVehicle) SetModel(e interface{}) {
	v.model = e
}

func (v *volkswagenVehicle) PrintDetails() string {
	return fmt.Sprintf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"+
		"~ Volkswagen %s\n"+
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", v.model)
}
