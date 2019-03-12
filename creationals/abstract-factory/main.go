package main

import (
	"fmt"

	"github.com/made2591/go-patterns/creationals/abstract-factory/consumer"
	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
)

func main() {
	c := consumer.NewConsumer(factory.NewCarFactory())
	fmt.Println(c.GetLicense().GetType())
	c = consumer.NewConsumer(factory.NewBusFactory())
	fmt.Println(c.GetLicense().GetType())
}
