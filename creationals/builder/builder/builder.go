package builder

import (
	"github.com/made2591/go-patterns/builder/vehicle"
)

type Builder interface {
	AddWheel()
	AddSteeringWheel()
	AddEngine()
	AddAirbag()
	GetVehicle() vehicle.Vehicle
}

type poloBuilder struct {
	polo vehicle.Vehicle
}

func NewPoloBuilder() Builder {
	return &poloBuilder{}
}

func (v *poloBuilder) AddWheel() {
	v.wheel = vehicle.VehicleStandardWheel
}

func (v *poloBuilder) AddSteeringWheel() {
	v.steeringWheel = vehicle.VehicleStandardSteeringWheel
}

func (v *poloBuilder) AddEngine() {
	v.engine = vehicle.VehicleStandardEngine
}

func (v *poloBuilder) AddAirbag() {
	v.airbagSystem = vehicle.VehicleStandardAirbagSystem
}

func (v *poloBuilder) GetVehicle() vehicle.Vehicle {
	return v.polo
}

type golfBuilder struct {
	golf vehicle.Vehicle
}

func NewGolfBuilder() Builder {
	return &golfBuilder{}
}

func (v *golfBuilder) AddWheel() {
	v.wheel = vehicle.VehicleSpecialWheel
}

func (v *golfBuilder) AddSteeringWheel() {
	v.steeringWheel = vehicle.VehicleSpecialSteeringWheel
}

func (v *golfBuilder) AddEngine() {
	v.engine = vehicle.VehicleSpecialEngine
}

func (v *golfBuilder) AddAirbag() {
	v.airbagSystem = vehicle.VehicleSpecialAirbagSystem
}

func (v *golfBuilder) GetVehicle() vehicle.Vehicle {
	return v.golf
}
