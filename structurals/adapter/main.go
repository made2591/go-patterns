package main

import (
	"fmt"

	"github.com/made2591/go-patterns/structurals/adapter/adapter"
	"github.com/made2591/go-patterns/structurals/adapter/reseller"
	"github.com/made2591/go-patterns/structurals/adapter/vehicle"
)

func main() {

	// init internal vehicle system with incompatible method for this application:
	internalVehicleSystem := vehicle.NewInternalVehicleSystem()

	// init the adapter that will manage the requests for the incompatible method
	adapterVehicleSystem := adapter.NewStoreAdapter(internalVehicleSystem)

	// init the reseller vehicle system by passing the adapter
	resellerVehicleSystem := reseller.NewResellerStore(adapterVehicleSystem)

	// ask for the vehicle in the requested format
	fmt.Println(resellerVehicleSystem.GetAvailableVehicles())

}
