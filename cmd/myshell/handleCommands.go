package main

import (
	"fmt"
	"harto.dev/trash/cmd/myshell/builtin"
	"os"
	"os/exec"
)

func HandleCommands(command string, args []string) {

	// Parse command and arguments
	isBuiltin := builtin.HandleBuiltins(command, args)
	if !isBuiltin {
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("%s: command not found\n", command)
			return
		}
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Printf("%s: error executing command\n", command)
		}
	}
}
