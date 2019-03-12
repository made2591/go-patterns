package license

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

	if strings.Compare(carTipe, LicenseCarType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", carTipe, LicenseCarType)
	}

	bus := NewBus()
	busTipe := bus.GetType().(string)

	if strings.Compare(busTipe, LicenseBusType) != 0 {
		t.Errorf("GetType was() incorrect, got %s, wanted %s", busTipe, LicenseBusType)
	}

}
