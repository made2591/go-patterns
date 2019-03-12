package vehicle

type Vehicle interface {
	GetModel() interface{}
	SetModel(m interface{})
}

type car struct {
	model interface{}
}

func New() Vehicle {
	return &car{}
}

func (c *car) GetModel() interface{} {
	return c.model
}

func (c *car) SetModel(m interface{}) {
	c.model = m
}
