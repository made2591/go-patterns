package consumer

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
	"github.com/made2591/go-patterns/creationals/abstract-factory/license"
	"github.com/made2591/go-patterns/creationals/abstract-factory/vehicle"
)

func TestNewConsumer(t *testing.T) {

	bus := factory.NewBusFactory()
	consumer := NewConsumer(bus)

	if consumer == nil {
		t.Errorf("NewConsumer was incorrect, got: nil")
	}

}

func TestGetVehicle(t *testing.T) {

	car := factory.NewCarFactory()
	consumer := NewConsumer(car)
	tipe := consumer.GetVehicle().GetTypeVehicle().(string)

	if strings.Compare(tipe, vehicle.VehicleCarType) != 0 {
		t.Errorf("GetTypeVehicle was incorrect, got %s, wanted %s", tipe, vehicle.VehicleCarType)
	}

	bus := factory.NewBusFactory()
	consumer = NewConsumer(bus)
	tipe = consumer.GetVehicle().GetTypeVehicle().(string)

	if strings.Compare(tipe, vehicle.VehicleBusType) != 0 {
		t.Errorf("GetTypeVehicle was incorrect, got %s, wanted %s", tipe, vehicle.VehicleBusType)
	}

}

func TestGetLicense(t *testing.T) {

	car := factory.NewCarFactory()
	consumer := NewConsumer(car)
	tipe := consumer.GetLicense().GetTypeLicense().(string)

	if strings.Compare(tipe, license.LicenseCarType) != 0 {
		t.Errorf("GetTypeLicense was incorrect, got %s, wanted %s", tipe, license.LicenseCarType)
	}

	bus := factory.NewBusFactory()
	consumer = NewConsumer(bus)
	tipe = consumer.GetLicense().GetTypeLicense().(string)

	if strings.Compare(tipe, license.LicenseBusType) != 0 {
		t.Errorf("GetTypeLicense was incorrect, got %s, wanted %s", tipe, license.LicenseBusType)
	}

}
