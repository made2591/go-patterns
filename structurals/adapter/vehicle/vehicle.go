package vehicle

const (
	VehicleCarType = "car"
	VehicleBusType = "bus"
	VehicleNDModel = "MODEL_NOT_DEFINED"
)

type Vehicle interface {
	GetTypeVehicle() interface{}
	GetModelVehicle() interface{}
	SetModelVehicle(m interface{})
}

type vehicle struct {
	tipe  interface{}
	model interface{}
}

func NewCarVehicle() Vehicle {
	return &vehicle{tipe: VehicleCarType, model: VehicleNDModel}
}

func NewBusVehicle() Vehicle {
	return &vehicle{tipe: VehicleBusType, model: VehicleNDModel}
}

func (v *vehicle) GetTypeVehicle() interface{} {
	return v.tipe
}

func (v *vehicle) GetModelVehicle() interface{} {
	return v.model
}

func (v *vehicle) SetModelVehicle(m interface{}) {
	v.model = m
}

type InternalVehicleSystem interface {
	GetAvailableVehicle() []Vehicle
}

type internalVehicleSystem struct{}

func NewInternalVehicleSystem() InternalVehicleSystem {
	return &internalVehicleSystem{}
}

func (ivs *internalVehicleSystem) GetAvailableVehicle() []Vehicle {

	vs := make([]Vehicle, 2)

	vs[0] = NewCarVehicle()
	vs[1] = NewBusVehicle()

	return vs

}
