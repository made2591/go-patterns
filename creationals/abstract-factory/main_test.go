package main

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/abstract-factory/consumer"
	"github.com/made2591/go-patterns/creationals/abstract-factory/factory"
)

func TestMain(t *testing.T) {

	f := factory.NewCarFactory()
	c := consumer.NewConsumer(f)

	if f == nil || c == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

}
