package vehicle

const (
	VehicleCarType = "car"
	VehicleBusType = "bus"
	VehicleNDModel = "MODEL_NOT_DEFINED"
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
	return &vehicle{tipe: VehicleCarType, model: VehicleNDModel}
}

func NewBus() Vehicle {
	return &vehicle{tipe: VehicleBusType, model: VehicleNDModel}
}

func (v *vehicle) GetType() interface{} {
	return v.tipe
}

func (v *vehicle) GetModel() interface{} {
	return v.model
}

func (v *vehicle) SetModel(m interface{}) {
	v.model = m
}
