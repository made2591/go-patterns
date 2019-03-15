package license

const (
	LicenseCarType = "B"
	LicenseBusType = "C"
)

type License interface {
	GetTypeLicense() interface{}
}

type license struct {
	tipe interface{}
}

func NewCarLicense() License {
	return &license{tipe: LicenseCarType}
}

func NewBusLicense() License {
	return &license{tipe: LicenseBusType}
}

func (v *license) GetTypeLicense() interface{} {
	return v.tipe
}
