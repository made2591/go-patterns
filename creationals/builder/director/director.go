package director

type Director interface {
	MakeVehicle(b builder.Builder)
}

type director struct {
}

func NewDirector() {
	return &director{}
}

func (d *director) MakeMeal(b builder.Builder) {
	b.AddWheel()
	b.AddSteeringWheel()
	b.AddEngine()
	b.AddAirbag()
}
