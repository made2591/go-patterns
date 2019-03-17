package main

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/builder/builder"
	"github.com/made2591/go-patterns/creationals/builder/director"
)

func main() {
	director := director.NewDirector()
	poloBuilder := builder.NewPoloBuilder()
	director.MakeMeal(poloBuilder)
	poloVehicle := poloBuilder.GetVehicle()
	fmt.Println(golfVehicle.GetWheel())
	fmt.Println(golfVehicle.GetEngine())

	golfBuilder := builder.NewGolfBuilder()
	director.MakeMeal(golfBuilder)
	golfVehicle := golfBuilder.GetVehicle()
	fmt.Println(golfVehicle.GetWheel())
	fmt.Println(golfVehicle.GetEngine())
}
