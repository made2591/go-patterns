package reseller

import (
	"github.com/made2591/go-patterns/structurals/adapter/adapter"
)

type Reseller interface {
	GetAvailableVehicles() string
}

type reseller struct {
	adapter adapter.Adapter
}

func NewResellerStore(adapter adapter.Adapter) Reseller {
	return &reseller{
		adapter: adapter,
	}
}

func (r *reseller) GetAvailableVehicles() string {
	return r.adapter.GetAvailableVehicles()
}
