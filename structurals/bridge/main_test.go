package main

import (
	"testing"

	"github.com/made2591/go-patterns/structurals/bridge/abstractinstaller"
	"github.com/made2591/go-patterns/structurals/bridge/systems"
)

func TestMain(t *testing.T) {

	// init window system
	windowSystem := systems.NewWindows()
	if windowSystem == nil {
		t.Errorf("Main was incorrect, got windowSystem: nil")
	}

	// init window system
	linuxSystem := systems.NewLinux()
	if linuxSystem == nil {
		t.Errorf("Main was incorrect, got linuxSystem: nil")
	}

	// init window system
	installerObject := abstractinstaller.NewInstaller()
	if installerObject == nil {
		t.Errorf("Main was incorrect, got installerObject: nil")
	}

}
