package abstractinstaller

import "github.com/made2591/go-patterns/structurals/bridge/abstractsystems"

type Installer interface {
	Setup(message string) string
	Set(abstractsystems abstractsystems.System)
	Get() abstractsystems.System
}

type installer struct {
	system abstractsystems.System
}

func NewInstaller() Installer {
	return &installer{}
}

func (i *installer) Set(s abstractsystems.System) {
	i.system = s
}

func (i *installer) Get() abstractsystems.System {
	return i.system
}

func (i *installer) Setup(message string) string {
	if i.system == nil {
		return "System not identified"
	}
	return i.system.Installer(message)
}
