package vehicle

import (
	"strings"
	"testing"
)

func TestNewVehicle(t *testing.T) {

	vehicle := NewVehicle()

	if vehicle == nil {
		t.Errorf("NewVehicle was incorrect, got: nil")
	}

}

func TestGetWheel(t *testing.T) {

	vehicle := NewVehicle()
	vehicleWheel := vehicle.GetWheel().(string)

	if strings.Compare(vehicleWheel, VehicleStandardWheel) != 0 {
		t.Errorf("GetWheel was incorrect, got %s, wanted %s", vehicleWheel, VehicleStandardWheel)
	}

}

func TestSetWheel(t *testing.T) {

	vehicle := NewVehicle()
	vehicle.SetWheel(VehicleSpecialWheel)
	vehicleWheel := vehicle.GetWheel().(string)

	if strings.Compare(vehicleWheel, VehicleSpecialWheel) != 0 {
		t.Errorf("SetWheel was incorrect, got %s, wanted %s", vehicleWheel, VehicleSpecialWheel)
	}

}

func TestGetSteeringWheel(t *testing.T) {

	vehicle := NewVehicle()
	vehicleSteeringWheel := vehicle.GetSteeringWheel().(string)

	if strings.Compare(vehicleSteeringWheel, VehicleStandardSteeringWheel) != 0 {
		t.Errorf("GetSteeringWheel was incorrect, got %s, wanted %s", vehicleSteeringWheel, VehicleStandardSteeringWheel)
	}

}

func TestSetSteeringWheel(t *testing.T) {

	vehicle := NewVehicle()
	vehicle.SetSteeringWheel(VehicleSpecialSteeringWheel)
	vehicleSteeringWheel := vehicle.GetSteeringWheel().(string)

	if strings.Compare(vehicleSteeringWheel, VehicleSpecialSteeringWheel) != 0 {
		t.Errorf("SetSteeringWheel was incorrect, got %s, wanted %s", vehicleSteeringWheel, VehicleSpecialSteeringWheel)
	}

}

func TestGetEngine(t *testing.T) {

	vehicle := NewVehicle()
	vehicleEngine := vehicle.GetEngine().(string)

	if strings.Compare(vehicleEngine, VehicleStandardEngine) != 0 {
		t.Errorf("GetEngine was incorrect, got %s, wanted %s", vehicleEngine, VehicleStandardEngine)
	}

}

func TestSetEngine(t *testing.T) {

	vehicle := NewVehicle()
	vehicle.SetEngine(VehicleSpecialEngine)
	vehicleEngine := vehicle.GetEngine().(string)

	if strings.Compare(vehicleEngine, VehicleSpecialEngine) != 0 {
		t.Errorf("SetEngine was incorrect, got %s, wanted %s", vehicleEngine, VehicleSpecialEngine)
	}

}

func TestGetAirbagSystem(t *testing.T) {

	vehicle := NewVehicle()
	vehicleAirbagSystem := vehicle.GetAirbagSystem().(string)

	if strings.Compare(vehicleAirbagSystem, VehicleStandardAirbagSystem) != 0 {
		t.Errorf("GetAirbagSystem was incorrect, got %s, wanted %s", vehicleAirbagSystem, VehicleStandardAirbagSystem)
	}

}

func TestSetAirbagSystem(t *testing.T) {

	vehicle := NewVehicle()
	vehicle.SetAirbagSystem(VehicleSpecialAirbagSystem)
	vehicleAirbagSystem := vehicle.GetAirbagSystem().(string)

	if strings.Compare(vehicleAirbagSystem, VehicleSpecialAirbagSystem) != 0 {
		t.Errorf("SetAirbagSystem was incorrect, got %s, wanted %s", vehicleAirbagSystem, VehicleSpecialAirbagSystem)
	}

}
