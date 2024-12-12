package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtins = []string{"exit", "echo", "type", "pwd", "cd"}

func main() {
	for i := 0; i < 3; i = i {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			os.Exit(1)
		}

		message := strings.TrimSpace(command)
		commands := strings.Split(message, " ")

		switch commands[0] {

		case "exit":
			switch commands[1] {
			case "0":
				os.Exit(0)
			default:
				os.Exit(0)
			}

		case "echo":
			i := 1
			for i < len(commands)-1 {
				fmt.Print(commands[i] + " ")
				i++
			}
			fmt.Println(commands[i])

		case "type":
			builtin_type(commands)

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)

		case "cd":
			os.Chdir(commands[1])
		default:
			command := exec.Command(commands[0], commands[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", commands[0])
			}
		}
	}
}

func builtin_type(commands []string) {
	if len(commands) < 2 {
		return
	} else {
		for _, b := range builtins {
			if b == commands[1] {
				fmt.Println(commands[1] + " is a shell builtin")
				return
			}

		}
		env := os.Getenv("PATH")
		path := strings.Split(env, ":")
		for _, p := range path {
			if _, err := os.Stat(p + "/" + commands[1]); err == nil {
				fmt.Println(commands[1] + " is " + p + "/" + commands[1])
				return
			}
		}
		fmt.Println(commands[1] + ": not found")
	}
}
