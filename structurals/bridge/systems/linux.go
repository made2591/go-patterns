package systems

import (
	"fmt"

	"github.com/made2591/go-patterns/structurals/bridge/abstractsystems"
)

type linux struct {
}

func NewLinux() abstractsystems.System {
	return &linux{}
}

func (l *linux) Installer(message string) string {
	return fmt.Sprintf("Installer for Linux start: %s", message)
}
