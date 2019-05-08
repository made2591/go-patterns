package systems

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewLinux(t *testing.T) {

	linux := NewLinux()

	if linux == nil {
		t.Errorf("NewLinux was incorrect, got: nil")
	}

}

func TestLInstaller(t *testing.T) {

	linux := NewLinux()
	message := "test"
	expected := fmt.Sprintf("Installer for Linux start: %s", message)
	computed := linux.Installer(message)

	if strings.Compare(expected, computed) != 0 {
		t.Errorf("Installer was incorrect, got: nil")
	}

}
