package director

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/builder/builder"
)

func TestNewDirector(t *testing.T) {

	director := NewDirector()

	if director == nil {
		t.Errorf("NewDirector was incorrect, got: nil")
	}

}

func TestMakeVehicle(t *testing.T) {

	director := NewDirector()
	poloBuilder := builder.NewPoloBuilder()
	director.MakeVehicle(poloBuilder)
	poloVehicle := poloBuilder.GetVehicle()
	if poloVehicle == nil {
		t.Errorf("MakeVehicle was incorrect, got: nil")
	}

}
