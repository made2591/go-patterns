package consumer

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

func TestNewConsumer(t *testing.T) {

	bus := factory.NewBusFactory()
	consumer := NewConsumer(bus)

	if consumer == nil {
		t.Errorf("NewConsumer was() incorrect, got: nil")
	}

}

func TestGetVehicle(t *testing.T) {

	car := factory.NewCarFactory()
	consumer := NewConsumer(car)
	tipe := consumer.GetVehicle().GetType().(string)

	if strings.Compare(tipe, vehicle.VehicleCarType) != 0 {
		t.Errorf("GetVehicle was() incorrect, got %s, wanted %s", tipe, vehicle.VehicleCarType)
	}

	bus := factory.NewBusFactory()
	consumer = NewConsumer(bus)
	tipe = consumer.GetVehicle().GetType().(string)

	if strings.Compare(tipe, vehicle.VehicleBusType) != 0 {
		t.Errorf("GetVehicle was() incorrect, got %s, wanted %s", tipe, vehicle.VehicleBusType)
	}

}
