package systems

import (
	"fmt"

	"github.com/made2591/go-patterns/structurals/bridge/abstractsystems"
)

type windows struct {
}

func NewWindows() abstractsystems.System {
	return &windows{}
}

func (w *windows) Installer(message string) string {
	return fmt.Sprintf("Installer for Windows start: %s", message)
}
