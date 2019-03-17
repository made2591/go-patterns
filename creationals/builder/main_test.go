package main

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/builder/builder"
	"github.com/made2591/go-patterns/creationals/builder/director"
)

func TestMain(t *testing.T) {

	director := director.NewDirector()

	if director == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	poloBuilder := builder.NewPoloBuilder()
	director.MakeVehicle(poloBuilder)
	poloVehicle := poloBuilder.GetVehicle()

	if poloBuilder == nil || poloVehicle == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	golfBuilder := builder.NewGolfBuilder()
	director.MakeVehicle(golfBuilder)
	golfVehicle := golfBuilder.GetVehicle()

	if golfBuilder == nil || golfVehicle == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

}
