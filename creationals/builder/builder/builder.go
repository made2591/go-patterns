package builder

import (
	"github.com/made2591/go-patterns/creationals/builder/vehicle"
)

type Builder interface {
	AddWheel()
	AddSteeringWheel()
	AddEngine()
	AddAirbagSystem()
	GetVehicle() vehicle.Vehicle
}

type poloBuilder struct {
	polo vehicle.Vehicle
}

func NewPoloBuilder() Builder {
	return &poloBuilder{polo: vehicle.NewVehicle()}
}

func (v *poloBuilder) AddWheel() {
	v.polo.SetWheel(vehicle.VehicleStandardWheel)
}

func (v *poloBuilder) AddSteeringWheel() {
	v.polo.SetSteeringWheel(vehicle.VehicleStandardSteeringWheel)
}

func (v *poloBuilder) AddEngine() {
	v.polo.SetEngine(vehicle.VehicleStandardEngine)
}

func (v *poloBuilder) AddAirbagSystem() {
	v.polo.SetAirbagSystem(vehicle.VehicleStandardAirbagSystem)
}

func (v *poloBuilder) GetVehicle() vehicle.Vehicle {
	return v.polo
}

type golfBuilder struct {
	golf vehicle.Vehicle
}

func NewGolfBuilder() Builder {
	return &golfBuilder{golf: vehicle.NewVehicle()}
}

func (v *golfBuilder) AddWheel() {
	v.golf.SetWheel(vehicle.VehicleSpecialWheel)
}

func (v *golfBuilder) AddSteeringWheel() {
	v.golf.SetSteeringWheel(vehicle.VehicleSpecialSteeringWheel)
}

func (v *golfBuilder) AddEngine() {
	v.golf.SetEngine(vehicle.VehicleSpecialEngine)
}

func (v *golfBuilder) AddAirbagSystem() {
	v.golf.SetAirbagSystem(vehicle.VehicleSpecialAirbagSystem)
}

func (v *golfBuilder) GetVehicle() vehicle.Vehicle {
	return v.golf
}
