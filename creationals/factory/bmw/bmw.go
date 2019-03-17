package bmw

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/factory/vehicle"
)

const (
	SerieOneModel = "Serie 1"
	SerieTwoModel = "Serie 2"
)

type bmwVehicle struct {
	model interface{}
}

func createBMWSerieOne() vehicle.Vehicle {
	return &bmwVehicle{
		model: SerieOneModel,
	}
}

func createBMWSerieTwo() vehicle.Vehicle {
	return &bmwVehicle{
		model: SerieTwoModel,
	}
}

func (v *bmwVehicle) GetModel() interface{} {
	return v.model
}

func (v *bmwVehicle) SetModel(e interface{}) {
	v.model = e
}

func (v *bmwVehicle) PrintDetails() string {
	return fmt.Sprintf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"+
		"~ BMW %s\n"+
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", v.model)
}
