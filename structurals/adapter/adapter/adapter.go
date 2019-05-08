package adapter

import (
	"bytes"
	"fmt"

	"github.com/made2591/go-patterns/structurals/adapter/vehicle"
)

type Adapter interface {
	GetAvailableVehicles() string
}

type adapter struct {
	internalVehicleSystem vehicle.InternalVehicleSystem
}

func NewStoreAdapter(ivs vehicle.InternalVehicleSystem) Adapter {
	return &adapter{
		internalVehicleSystem: ivs,
	}
}

func (r *adapter) GetAvailableVehicles() string {
	var buffer bytes.Buffer
	for _, vehicle := range r.internalVehicleSystem.GetAvailableVehicles() {
		buffer.WriteString(fmt.Sprintf("%s, %s\n", vehicle.GetTypeVehicle().(string), vehicle.GetModelVehicle().(string)))
	}
	return buffer.String()
}
