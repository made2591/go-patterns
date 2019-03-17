package vehicle

const (
	VehicleStandardWheel         = "standard wheel"
	VehicleStandardSteeringWheel = "standard steering wheel"
	VehicleStandardEngine        = "standard engine"
	VehicleStandardAirbagSystem  = "standard airbags"
	VehicleSpecialWheel          = "special wheel"
	VehicleSpecialSteeringWheel  = "special steering wheel"
	VehicleSpecialEngine         = "special engine"
	VehicleSpecialAirbagSystem   = "special airbags"
)

type Vehicle interface {
	GetWheel() interface{}
	GetSteeringWheel() interface{}
	GetEngine() interface{}
	GetAirbagSystem() interface{}
	SetWheel(a interface{})
	SetSteeringWheel(a interface{})
	SetEngine(a interface{})
	SetAirbagSystem(a interface{})
}

type vehicle struct {
	wheel         interface{}
	steeringWheel interface{}
	engine        interface{}
	airbagSystem  interface{}
}

func NewVehicle() Vehicle {
	return &vehicle{
		wheel:         VehicleStandardWheel,
		steeringWheel: VehicleStandardSteeringWheel,
		engine:        VehicleStandardEngine,
		airbagSystem:  VehicleStandardAirbagSystem,
	}
}

func (v *vehicle) GetWheel() interface{} {
	return v.wheel
}

func (v *vehicle) GetSteeringWheel() interface{} {
	return v.steeringWheel
}

func (v *vehicle) GetEngine() interface{} {
	return v.engine
}

func (v *vehicle) GetAirbagSystem() interface{} {
	return v.airbagSystem
}

func (v *vehicle) SetWheel(w interface{}) {
	v.wheel = w
}

func (v *vehicle) SetSteeringWheel(sw interface{}) {
	v.steeringWheel = sw
}

func (v *vehicle) SetEngine(e interface{}) {
	v.engine = e
}

func (v *vehicle) SetAirbagSystem(as interface{}) {
	v.airbagSystem = as
}
