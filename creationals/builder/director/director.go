package director

import "github.com/made2591/go-patterns/creationals/builder/builder"

type Director interface {
	MakeVehicle(b builder.Builder)
}

type director struct {
}

func NewDirector() Director {
	return &director{}
}

func (d *director) MakeVehicle(b builder.Builder) {
	b.AddWheel()
	b.AddSteeringWheel()
	b.AddEngine()
	b.AddAirbagSystem()
}
