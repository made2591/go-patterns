package license

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

func TestNewCar(t *testing.T) {

	car := NewCar()

	if car == nil {
		t.Errorf("NewCar was() incorrect, got: nil")
	}

}

func TestNewBus(t *testing.T) {

	car := NewCar()

	if car == nil {
		t.Errorf("NewCar was() incorrect, got: nil")
	}

}

func TestGetType(t *testing.T) {

	car := NewCar()
	carTipe := car.GetType().(string)

	if strings.Compare(carTipe, vehicle.VehicleCarType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", carTipe, vehicle.VehicleCarType)
	}

	bus := NewBus()
	busTipe := bus.GetType().(string)

	if strings.Compare(busTipe, vehicle.VehicleBusType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", busTipe, vehicle.VehicleBusType)
	}

}
