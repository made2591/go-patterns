package bmw

import "testing"

func TestNewBMWFactory(t *testing.T) {

	factory := NewBmwFactory()

	if factory == nil {
		t.Errorf("NewBmwFactory was incorrect, got: nil")
	}

}
