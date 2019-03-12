package license

const (
	licenseCarType = "car"
	licenseBusType = "bus"
)

type License interface {
	GetType() interface{}
}

type license struct {
	tipe interface{}
}

func NewCar() License {
	return &license{tipe: licenseCarType}
}

func NewBus() License {
	return &license{tipe: licenseBusType}
}

func (v *license) GetType() interface{} {
	return v.tipe
}
