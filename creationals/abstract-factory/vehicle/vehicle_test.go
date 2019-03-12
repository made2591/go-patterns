package vehicle

import (
	"strings"
	"testing"
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

	if strings.Compare(carTipe, VehicleCarType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", carTipe, VehicleCarType)
	}

	bus := NewBus()
	busTipe := bus.GetType().(string)

	if strings.Compare(busTipe, VehicleBusType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", busTipe, VehicleBusType)
	}

}

func TestGetModel(t *testing.T) {

	car := NewCar()
	carModel := car.GetModel().(string)

	if strings.Compare(carModel, VehicleNDModel) != 0 {
		t.Errorf("GetModel was() incorrect, got %s, wanted %s", carModel, VehicleNDModel)
	}

	bus := NewBus()
	busModel := bus.GetModel().(string)

	if strings.Compare(busModel, VehicleNDModel) != 0 {
		t.Errorf("GetModel was() incorrect, got %s, wanted %s", busModel, VehicleNDModel)
	}

}

func TestSetModel(t *testing.T) {

	car := NewCar()
	car.SetModel("Punto")
	carModel := car.GetModel().(string)

	if strings.Compare(carModel, "Punto") != 0 {
		t.Errorf("SetModel was() incorrect, got %s, wanted %s", carModel, "Punto")
	}

	bus := NewBus()
	bus.SetModel("Iveco")
	busModel := bus.GetModel().(string)

	if strings.Compare(busModel, "Iveco") != 0 {
		t.Errorf("SetModel was() incorrect, got %s, wanted %s", busModel, "Iveco")
	}

}
