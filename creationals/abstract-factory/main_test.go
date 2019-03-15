package main

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/consumer"
	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
)

func Testmain(t *testing.T) {

	fc := factory.NewCarFactory()
	c := consumer.NewConsumer(fc)

	if fc == nil || c == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

	fb := factory.NewBusFactory()
	b := consumer.NewConsumer(fb)

	if fb == nil || b == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

}
