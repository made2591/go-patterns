package license

const (
	LicenseCarType = "B"
	LicenseBusType = "C"
)

type License interface {
	GetType() interface{}
}

type license struct {
	tipe interface{}
}

func NewCar() License {
	return &license{tipe: LicenseCarType}
}

func NewBus() License {
	return &license{tipe: LicenseBusType}
}

func (v *license) GetType() interface{} {
	return v.tipe
}
