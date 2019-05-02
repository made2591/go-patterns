package logger

import (
	"testing"
)

func TestNew(t *testing.T) {

	logger := New()

	if logger == nil {
		t.Errorf("New was incorrect, got: nil")
	}

}

func TestGetLevel(t *testing.T) {

	logger := New()
	logger.SetLevel(2)
	level := logger.GetLevel()

	if level != 2 {
		t.Errorf("GetLevel was incorrect, got %d, wanted %d", logger, level)
	}

}

func TestSetLevel(t *testing.T) {

	logger := New()
	logger.SetLevel(3)

	if level := logger.GetLevel(); level != 3 {
		t.Errorf("SetLevel was incorrect, got %d, wanted %d", logger, level)
	}

}
