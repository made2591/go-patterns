package vehicle

import (
	"strings"
	"testing"
)

func TestNewCarVehicle(t *testing.T) {

	car := NewCarVehicle()

	if car == nil {
		t.Errorf("NewCarVehicle was incorrect, got: nil")
	}

}

func TestNewBusVehicle(t *testing.T) {

	bus := NewBusVehicle()

	if bus == nil {
		t.Errorf("NewBusVehicle was incorrect, got: nil")
	}

}

func TestGetTypeVehicle(t *testing.T) {

	car := NewCarVehicle()
	carTipe := car.GetTypeVehicle().(string)

	if strings.Compare(carTipe, VehicleCarType) != 0 {
		t.Errorf("GetTypeVehicle was incorrect, got %s, wanted %s", carTipe, VehicleCarType)
	}

	bus := NewBusVehicle()
	busTipe := bus.GetTypeVehicle().(string)

	if strings.Compare(busTipe, VehicleBusType) != 0 {
		t.Errorf("GetTypeVehicle was incorrect, got %s, wanted %s", busTipe, VehicleBusType)
	}

}

func TestGetModelVehicle(t *testing.T) {

	car := NewCarVehicle()
	carModel := car.GetModelVehicle().(string)

	if strings.Compare(carModel, VehicleNDModel) != 0 {
		t.Errorf("GetModelVehicle was incorrect, got %s, wanted %s", carModel, VehicleNDModel)
	}

	bus := NewBusVehicle()
	busModel := bus.GetModelVehicle().(string)

	if strings.Compare(busModel, VehicleNDModel) != 0 {
		t.Errorf("GetModelVehicle was incorrect, got %s, wanted %s", busModel, VehicleNDModel)
	}

}

func TestSetModelVehicle(t *testing.T) {

	car := NewCarVehicle()
	car.SetModelVehicle("Punto")
	carModel := car.GetModelVehicle().(string)

	if strings.Compare(carModel, "Punto") != 0 {
		t.Errorf("SetModelVehicle was incorrect, got %s, wanted %s", carModel, "Punto")
	}

	bus := NewBusVehicle()
	bus.SetModelVehicle("Iveco")
	busModel := bus.GetModelVehicle().(string)

	if strings.Compare(busModel, "Iveco") != 0 {
		t.Errorf("SetModelVehicle was incorrect, got %s, wanted %s", busModel, "Iveco")
	}

}
