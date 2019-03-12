package vehicle

const (
	vehicleCarType = "car"
	vehicleBusType = "bus"
)

type Vehicle interface {
	GetType() interface{}
	GetModel() interface{}
	SetModel(m interface{})
}

type vehicle struct {
	tipe  interface{}
	model interface{}
}

func NewCar() Vehicle {
	return &vehicle{tipe: vehicleCarType}
}

func NewBus() Vehicle {
	return &vehicle{tipe: vehicleBusType}
}

func (v *vehicle) GetType() interface{} {
	return v.tipe
}

func (v *vehicle) GetModel() interface{} {
	return v.model
}

func (v *vehicle) SetModel(m interface{}) {
	v.tipe = m
}
