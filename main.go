package main

import (
	"mai/command"
	"os"
)

func main() {
	registry := command.InitCommandRegistry()
	registry.LoadCommands()

	if len(os.Args) < 2 {
		registry.Run("help", nil)

	} else {
		cmdName := os.Args[1]
		cmdArgs := os.Args[2:]
		registry.Run(cmdName, cmdArgs)

	}
}
