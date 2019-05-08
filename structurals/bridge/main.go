package main

import (
	"fmt"

	"github.com/made2591/go-patterns/structurals/bridge/abstractinstaller"
	"github.com/made2591/go-patterns/structurals/bridge/systems"
)

func main() {

	windowSystem := systems.NewWindows()
	linuxSystem := systems.NewLinux()

	installerObject := abstractinstaller.NewInstaller()
	installerObject.Set(windowSystem)
	fmt.Println(installerObject.Setup("Start installer"))

	installerObject.Set(linuxSystem)
	fmt.Println(installerObject.Setup("Start installer"))

}
