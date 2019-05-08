package abstractinstaller

import (
	"testing"

	"github.com/made2591/go-patterns/structurals/bridge/systems"
)

func TestNewInstaller(t *testing.T) {

	installer := NewInstaller()

	if installer == nil {
		t.Errorf("NewInstaller was incorrect, got: nil")
	}

}

func TestGet(t *testing.T) {

	windowSystem := systems.NewWindows()
	installer := NewInstaller()
	installer.Set(windowSystem)

	if installer.Get() == nil {
		t.Errorf("Get was incorrect, got: nil")
	}

}

func TestSet(t *testing.T) {

	windowSystem := systems.NewWindows()
	installer := NewInstaller()
	installer.Set(windowSystem)

	if installer.Get() == nil {
		t.Errorf("Set was incorrect, got: nil")
	}

}
