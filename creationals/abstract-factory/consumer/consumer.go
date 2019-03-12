package consumer

import (
	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
	"github.com/made2591/go-patterns/creationals/abstract-factory/license"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

type Consumer interface {
	GetVehicle() vehicle.Vehicle
	GetLicense() license.License
}

type consumer struct {
	vehicle vehicle.Vehicle
	license license.License
}

func NewConsumer(factory factory.Factory) Consumer {
	return &consumer{vehicle: factory.NewVehicle(), license: factory.NewLicense()}
}

func (c *consumer) GetVehicle() vehicle.Vehicle {
	return c.vehicle
}

func (c *consumer) GetLicense() license.License {
	return c.license
}

func main() {
	c := NewConsumer(factory.NewFactory())
	c.GetLicense().GetType()
}
