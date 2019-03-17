package main

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/builder/builder"
	"github.com/made2591/go-patterns/creationals/builder/director"
)

func main() {
	director := director.NewDirector()

	poloBuilder := builder.NewPoloBuilder()
	director.MakeVehicle(poloBuilder)
	poloVehicle := poloBuilder.GetVehicle()
	fmt.Println(poloVehicle.PrintDetails())

	golfBuilder := builder.NewGolfBuilder()
	director.MakeVehicle(golfBuilder)
	golfVehicle := golfBuilder.GetVehicle()
	fmt.Println(golfVehicle.PrintDetails())
}
