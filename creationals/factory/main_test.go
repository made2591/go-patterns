package main

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/factory/bmw"
	"github.com/made2591/go-patterns/creationals/factory/volkswagen"
)

func TestMain(t *testing.T) {

	volkswagenFactory := volkswagen.NewVolkswagenFactory()

	if volkswagenFactory == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	poloVehicle, err := volkswagenFactory.CreateVehicle(uint(1))

	if poloVehicle == nil || err != nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	golfVehicle, err := volkswagenFactory.CreateVehicle(uint(2))

	if golfVehicle == nil || err != nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	_, err = volkswagenFactory.CreateVehicle(uint(3))

	if err == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	bmwFactory := bmw.NewBmwFactory()

	if bmwFactory == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	serieOneVehicle, err := bmwFactory.CreateVehicle(uint(1))

	if serieOneVehicle == nil || err != nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	serieTwoVehicle, err := bmwFactory.CreateVehicle(uint(2))

	if serieTwoVehicle == nil || err != nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	_, err = bmwFactory.CreateVehicle(uint(3))

	if err == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

}
