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
	return vehicle.NewCarVehicle()
}

func (f *carFactory) NewLicense() license.License {
	return license.NewCarLicense()
}

type busFactory struct {
}

func NewBusFactory() Factory {
	return &busFactory{}
}

func (f *busFactory) NewVehicle() vehicle.Vehicle {
	return vehicle.NewBusVehicle()
}

func (f *busFactory) NewLicense() license.License {
	return license.NewBusLicense()
}
