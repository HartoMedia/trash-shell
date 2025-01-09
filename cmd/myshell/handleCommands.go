package main

import (
	"fmt"
	"harto.dev/trash/cmd/myshell/builtin"
	"os"
	"os/exec"
)

func handleCommands(input string) {
	// Parse command and arguments
	command, args := parseCommand(input)

	switch command {
	case "exit":
		builtin.Exit(args)
	case "echo":
		builtin.Echo(args)
	case "type":
		builtin.Type(args)
	case "cls":
		builtin.Cls()
	case "pwd":
		builtin.Pwd()
	case "cd":
		builtin.Cd(args)
	case "thx":
		builtin.Thx()
	default:
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("%s: command not found\n", command)
			break
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
