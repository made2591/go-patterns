package factory

import (
	"github.com/made2591/go-patterns/creationals/abstract-factory/license"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

type Factory interface {
	NewVehicle() vehicle.Vehicle
	NewLicense() license.License
}

type factory struct {
}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) NewVehicle() vehicle.Vehicle {
	return vehicle.New()
}

func (f *factory) NewLicense() license.License {
	return license.New()
}
