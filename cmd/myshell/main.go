package main

import (
	"bufio"
	"fmt"
	"harto.dev/trash/cmd/myshell/builtin"
	"os"
	"strings"
)

func main() {
	for {
		Ps1()
		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, "error reading input:", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}

		// Remove newline characters weil wichtig
		userInput = strings.Trim(userInput, "\r\n")

		// Skip empty lines
		if len(userInput) == 0 {
			continue
		}

		// Parse command and arguments from user input
		command, args := ParseCommand(userInput)

		// Handle da shit
		builtin.HandleCommands(command, args)
	}
}
