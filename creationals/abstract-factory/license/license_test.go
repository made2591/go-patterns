package license

import (
	"strings"
	"testing"
)

func TestNewCarLicense(t *testing.T) {

	car := NewCarLicense()

	if car == nil {
		t.Errorf("NewCarLicense was incorrect, got: nil")
	}

}

func TestNewBusLicense(t *testing.T) {

	bus := NewBusLicense()

	if bus == nil {
		t.Errorf("NewBusLicense was incorrect, got: nil")
	}

}

func TestGetType(t *testing.T) {

	car := NewCarLicense()
	carTipe := car.GetTypeLicense().(string)

	if strings.Compare(carTipe, LicenseCarType) != 0 {
		t.Errorf("GetTypeLicense was incorrect, got %s, wanted %s", carTipe, LicenseCarType)
	}

	bus := NewBusLicense()
	busTipe := bus.GetTypeLicense().(string)

	if strings.Compare(busTipe, LicenseBusType) != 0 {
		t.Errorf("GetTypeLicense was incorrect, got %s, wanted %s", busTipe, LicenseBusType)
	}

}
