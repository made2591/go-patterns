package builder

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/creationals/builder/vehicle"
)

func TestNewPoloBuilder(t *testing.T) {

	poloBuilder := NewPoloBuilder()

	if poloBuilder == nil {
		t.Errorf("NewPoloBuilder was incorrect, got: nil")
	}

}

func TestNewGolfBuilder(t *testing.T) {

	golfBuilder := NewGolfBuilder()

	if golfBuilder == nil {
		t.Errorf("NewGolfBuilder was incorrect, got: nil")
	}

}

func TestAddWheel(t *testing.T) {

	poloBuilder := NewPoloBuilder()
	poloBuilder.AddWheel()
	poloWheel := poloBuilder.GetVehicle().GetWheel().(string)

	if strings.Compare(poloWheel, vehicle.VehicleStandardWheel) != 0 {
		t.Errorf("AddWheel was incorrect, got %s, wanted %s", poloWheel, vehicle.VehicleStandardWheel)
	}

	golfBuilder := NewGolfBuilder()
	golfBuilder.AddWheel()
	golfWheel := golfBuilder.GetVehicle().GetWheel().(string)

	if strings.Compare(golfWheel, vehicle.VehicleSpecialWheel) != 0 {
		t.Errorf("AddWheel was incorrect, got %s, wanted %s", golfWheel, vehicle.VehicleSpecialWheel)
	}

}

func TestAddSteeringWheel(t *testing.T) {

	poloBuilder := NewPoloBuilder()
	poloBuilder.AddSteeringWheel()
	poloSteeringWheel := poloBuilder.GetVehicle().GetSteeringWheel().(string)

	if strings.Compare(poloSteeringWheel, vehicle.VehicleStandardSteeringWheel) != 0 {
		t.Errorf("AddSteeringWheel was incorrect, got %s, wanted %s", poloSteeringWheel, vehicle.VehicleStandardSteeringWheel)
	}

	golfBuilder := NewGolfBuilder()
	golfBuilder.AddSteeringWheel()
	golfSteeringWheel := golfBuilder.GetVehicle().GetSteeringWheel().(string)

	if strings.Compare(golfSteeringWheel, vehicle.VehicleSpecialSteeringWheel) != 0 {
		t.Errorf("AddSteeringWheel was incorrect, got %s, wanted %s", golfSteeringWheel, vehicle.VehicleSpecialSteeringWheel)
	}

}

func TestAddEngine(t *testing.T) {

	poloBuilder := NewPoloBuilder()
	poloBuilder.AddEngine()
	poloEngine := poloBuilder.GetVehicle().GetEngine().(string)

	if strings.Compare(poloEngine, vehicle.VehicleStandardEngine) != 0 {
		t.Errorf("AddEngine was incorrect, got %s, wanted %s", poloEngine, vehicle.VehicleStandardEngine)
	}

	golfBuilder := NewGolfBuilder()
	golfBuilder.AddEngine()
	golfEngine := golfBuilder.GetVehicle().GetEngine().(string)

	if strings.Compare(golfEngine, vehicle.VehicleSpecialEngine) != 0 {
		t.Errorf("AddEngine was incorrect, got %s, wanted %s", golfEngine, vehicle.VehicleSpecialEngine)
	}

}

func TestAddAirbagSystem(t *testing.T) {

	poloBuilder := NewPoloBuilder()
	poloBuilder.AddAirbagSystem()
	poloAirbagSystem := poloBuilder.GetVehicle().GetAirbagSystem().(string)

	if strings.Compare(poloAirbagSystem, vehicle.VehicleStandardAirbagSystem) != 0 {
		t.Errorf("AddAirbagSystem was incorrect, got %s, wanted %s", poloAirbagSystem, vehicle.VehicleStandardAirbagSystem)
	}

	golfBuilder := NewGolfBuilder()
	golfBuilder.AddAirbagSystem()
	golfAirbagSystem := golfBuilder.GetVehicle().GetAirbagSystem().(string)

	if strings.Compare(golfAirbagSystem, vehicle.VehicleSpecialAirbagSystem) != 0 {
		t.Errorf("AddAirbagSystem was incorrect, got %s, wanted %s", golfAirbagSystem, vehicle.VehicleSpecialAirbagSystem)
	}

}

func TestGetVehicle(t *testing.T) {

	poloBuilder := NewPoloBuilder()
	poloVehicle := poloBuilder.GetVehicle()

	if poloVehicle == nil {
		t.Errorf("GetVehicle was incorrect, got: nil")
	}

	golfBuilder := NewGolfBuilder()
	golfVehicle := golfBuilder.GetVehicle()

	if golfVehicle == nil {
		t.Errorf("GetVehicle was incorrect, got: nil")
	}

}
