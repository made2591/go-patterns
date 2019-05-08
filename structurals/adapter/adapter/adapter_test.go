package adapter

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/structurals/adapter/vehicle"
)

func TestNewStoreAdapter(t *testing.T) {

	internalVehicleSystem := vehicle.NewInternalVehicleSystem()
	storeAdapter := NewStoreAdapter(internalVehicleSystem)

	if storeAdapter == nil {
		t.Errorf("NewStoreAdapter was incorrect, got: nil")
	}

}

func TestGetAvailableVehicles(t *testing.T) {

	internalVehicleSystem := vehicle.NewInternalVehicleSystem()
	storeAdapter := NewStoreAdapter(internalVehicleSystem)
	wanted :=
		`car, Golf
bus, Iveco
`
	retrieved := storeAdapter.GetAvailableVehicles()
	if strings.Compare(retrieved, wanted) != 0 {
		t.Errorf("GetAvailableVehicles was incorrect, wanted %s got: %s", wanted, retrieved)
	}

}
