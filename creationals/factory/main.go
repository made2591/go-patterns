package main

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/factory/bmw"
	"github.com/made2591/go-patterns/creationals/factory/volkswagen"
)

func main() {
	volkswagenFactory := volkswagen.NewVolkswagenFactory()
	poloVehicle, err := volkswagenFactory.CreateVehicle(uint(1))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(poloVehicle.PrintDetails())
	}
	golfVehicle, err := volkswagenFactory.CreateVehicle(uint(2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(golfVehicle.PrintDetails())
	}
	_, err = volkswagenFactory.CreateVehicle(uint(2))
	if err != nil {
		fmt.Println(err)
	}

	bmwFactory := bmw.NewBmwFactory()
	serieOneVehicle, err := bmwFactory.CreateVehicle(uint(1))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serieOneVehicle.PrintDetails())
	}
	serieTwoVehicle, err := bmwFactory.CreateVehicle(uint(2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serieTwoVehicle.PrintDetails())
	}
	_, err = bmwFactory.CreateVehicle(uint(2))
	if err != nil {
		fmt.Println(err)
	}

}
