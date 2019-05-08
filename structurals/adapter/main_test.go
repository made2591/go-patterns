package main

import (
	"testing"

	"github.com/made2591/go-patterns/structurals/adapter/adapter"
	"github.com/made2591/go-patterns/structurals/adapter/reseller"
	"github.com/made2591/go-patterns/structurals/adapter/vehicle"
)

func TestMain(t *testing.T) {

	// init internal vehicle system with incompatible method for this application:
	internalVehicleSystem := vehicle.NewInternalVehicleSystem()
	if internalVehicleSystem == nil {
		t.Errorf("Main was incorrect, got internalVehicleSystem: nil")
	}

	// init the adapter that will manage the requests for the incompatible method
	adapterVehicleSystem := adapter.NewStoreAdapter(internalVehicleSystem)
	if adapterVehicleSystem == nil {
		t.Errorf("Main was incorrect, got adapterVehicleSystem: nil")
	}

	// init the reseller vehicle system by passing the adapter
	resellerVehicleSystem := reseller.NewResellerStore(adapterVehicleSystem)
	if resellerVehicleSystem == nil {
		t.Errorf("Main was incorrect, got resellerVehicleSystem: nil")
	}

}
