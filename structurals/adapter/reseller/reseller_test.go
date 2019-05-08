package reseller

import (
	"strings"
	"testing"

	"github.com/made2591/go-patterns/structurals/adapter/adapter"
	"github.com/made2591/go-patterns/structurals/adapter/vehicle"
)

func TestNewResellerStore(t *testing.T) {

	internalVehicleSystem := vehicle.NewInternalVehicleSystem()
	adapterVehicleSystem := adapter.NewStoreAdapter(internalVehicleSystem)
	reseller := NewResellerStore(adapterVehicleSystem)

	if reseller == nil {
		t.Errorf("NewResellerStore was incorrect, got: nil")
	}

}

func TestGetAvailableVehicles(t *testing.T) {

	internalVehicleSystem := vehicle.NewInternalVehicleSystem()
	adapterVehicleSystem := adapter.NewStoreAdapter(internalVehicleSystem)
	reseller := NewResellerStore(adapterVehicleSystem)
	wanted :=
		`car, Golf
bus, Iveco
`
	retrieved := reseller.GetAvailableVehicles()
	if strings.Compare(retrieved, wanted) != 0 {
		t.Errorf("GetAvailableVehicles was incorrect, wanted %s got: %s", wanted, retrieved)
	}

}
