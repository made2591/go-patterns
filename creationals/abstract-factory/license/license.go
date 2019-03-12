package license

type License interface {
	GetType() interface{}
	SetType(t interface{})
}

type b struct {
	tipe interface{}
}

func New() License {
	return &b{}
}

func (l *b) SetType(t interface{}) {
	l.tipe = t
}

func (l *b) GetType() interface{} {
	return l.tipe
}
