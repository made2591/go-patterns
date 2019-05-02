package main

import (
	"testing"

	"github.com/made2591/go-patterns/creationals/singleton/logger"
)

func TestMain(t *testing.T) {

	l1 := logger.New()
	l1.SetLevel(2)
	l2 := logger.New()
	l2.SetLevel(2)

	if l1 == nil || l2 == nil {
		t.Errorf("Main was incorrect, got: nil")
	}
	l1v := l1.GetLevel()
	l2v := l2.GetLevel()

	if l1v != l2v {
		t.Errorf("Main was incorrect, got: %d, %d", l1v, l2v)
	}

}
