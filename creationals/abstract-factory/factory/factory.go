package factory

import (
	"github.com/made2591/go-patterns/creationals/abstract-factory/license"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

type Factory interface {
	NewVehicle() vehicle.Vehicle
	NewLicense() license.License
}

type carFactory struct {
}

func NewCarFactory() Factory {
	return &carFactory{}
}

func (f *carFactory) NewVehicle() vehicle.Vehicle {
	return vehicle.NewCar()
}

func (f *carFactory) NewLicense() license.License {
	return license.NewCar()
}

type busFactory struct {
}

func NewBusFactory() Factory {
	return &busFactory{}
}

func (f *busFactory) NewVehicle() vehicle.Vehicle {
	return vehicle.NewBus()
}

func (f *busFactory) NewLicense() license.License {
	return license.NewBus()
}
