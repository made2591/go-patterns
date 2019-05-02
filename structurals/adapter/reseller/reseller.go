package reseller

type Reseller interface {
	GetAvailableVehicle() string
}

type reseller struct {
}
