package systems

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewWindows(t *testing.T) {

	linux := NewWindows()

	if linux == nil {
		t.Errorf("NewLinux was incorrect, got: nil")
	}

}

func TestWInstaller(t *testing.T) {

	linux := NewWindows()
	message := "test"
	expected := fmt.Sprintf("Installer for Windows start: %s", message)
	computed := linux.Installer(message)

	if strings.Compare(expected, computed) != 0 {
		t.Errorf("Installer was incorrect, got: nil")
	}

}
