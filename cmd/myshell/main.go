package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

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
			for i := 1; i < len(commands)-1; i++ {
				fmt.Print(commands[i] + " ")
			}
			fmt.Println(commands[i])
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
