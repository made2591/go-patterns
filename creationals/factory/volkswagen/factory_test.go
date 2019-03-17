package volkswagen

import "testing"

func TestNewVolkswagenFactory(t *testing.T) {

	factory := NewVolkswagenFactory()

	if factory == nil {
		t.Errorf("NewVolkswagenFactory was incorrect, got: nil")
	}

}

func TestCreateVehicle(t *testing.T) {

	volkswagenFactory := NewVolkswagenFactory()

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

}
