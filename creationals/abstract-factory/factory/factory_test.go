package factory

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/license"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

func TestNewCarFactory(t *testing.T) {

	car := NewCarFactory()

	if car == nil {
		t.Errorf("NewCarFactory was() incorrect, got: nil")
	}

}

func TestNewBusFactory(t *testing.T) {

	bus := NewBusFactory()

	if bus == nil {
		t.Errorf("NewBusFactory was() incorrect, got: nil")
	}

}

func TestNewVehicle(t *testing.T) {

	factoryCar := NewCarFactory()
	carVehicle := factoryCar.NewVehicle()
	vehicleTipe := carVehicle.GetType().(string)

	if strings.Compare(vehicleTipe, vehicle.VehicleCarType) != 0 {
		t.Errorf("NewVehicle was() incorrect, got %s, wanted %s", vehicleTipe, vehicle.VehicleCarType)
	}

	factoryBus := NewBusFactory()
	busVehicle := factoryBus.NewVehicle()
	vehicleTipe = busVehicle.GetType().(string)

	if strings.Compare(vehicleTipe, vehicle.VehicleBusType) != 0 {
		t.Errorf("NewVehicle was() incorrect, got %s, wanted %s", vehicleTipe, vehicle.VehicleBusType)
	}

}

func TestNewLicense(t *testing.T) {

	factoryCar := NewCarFactory()
	carVehicle := factoryCar.NewVehicle()
	licenseTipe := carVehicle.GetType().(string)

	if strings.Compare(licenseTipe, license.LicenseCarType) != 0 {
		t.Errorf("NewLicense was() incorrect, got %s, wanted %s", licenseTipe, license.LicenseCarType)
	}

	factoryBus := NewBusFactory()
	busLicense := factoryBus.NewLicense()
	licenseTipe = busLicense.GetType().(string)

	if strings.Compare(licenseTipe, license.LicenseBusType) != 0 {
		t.Errorf("NewLicense was() incorrect, got %s, wanted %s", licenseTipe, license.LicenseBusType)
	}

}
