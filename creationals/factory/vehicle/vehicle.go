package vehicle

type Vehicle interface {
	GetModel() interface{}
	SetModel(a interface{})
	PrintDetails() string
}
