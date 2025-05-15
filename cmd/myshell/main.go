package main

import (
	"bufio"
	"fmt"
	"harto.dev/trash/cmd/myshell/builtin"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var lastExitCode = 0

func main() {
	// Check if we're being called to execute a script
	if len(os.Args) > 1 {
		executeScript(os.Args[1], os.Args[2:])
		return
	}

	// Interactive mode
	for {
		// Display prompt
		displayPrompt()

		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, "error reading input:", err)
			if err != nil {
				os.Exit(1)
			}
			os.Exit(1)
		}

		// Remove newline characters
		userInput = strings.Trim(userInput, "\r\n")

		// Skip empty lines
		if len(userInput) == 0 {
			continue
		}

		// Parse command and arguments from user input
		command, args := ParseCommand(userInput)

		// Execute the command
		lastExitCode = builtin.HandleCommands(command, args)
	}
}

func displayPrompt() {
	// Check if PS1 environment variable is set
	ps1 := os.Getenv("PS1")
	if ps1 != "" {
		fmt.Print(ps1)
	} else {
		// Fall back to default prompt
		Ps1()
	}
}

func executeScript(scriptPath string, args []string) {
	// Check if file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", scriptPath)
		os.Exit(127)
	}

	// Read the file
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading script: %v\n", err)
		os.Exit(1)
	}

	// Check for shebang
	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "#!") {
		// Extract interpreter
		interpreter := strings.TrimPrefix(lines[0], "#!")
		interpreter = strings.TrimSpace(interpreter)

		// If the interpreter is this shell, execute the script line by line
		if filepath.Base(interpreter) == "myshell" || strings.HasSuffix(interpreter, "/myshell") {
			executeScriptLines(lines[1:])
		} else {
			// Otherwise, let the specified interpreter handle it
			cmd := exec.Command(interpreter, scriptPath)
			cmd.Args = append(cmd.Args, args...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					os.Exit(exitErr.ExitCode())
				}
				os.Exit(1)
			}
		}
	} else {
		// No shebang, execute the script line by line
		executeScriptLines(lines)
	}
}

func executeScriptLines(lines []string) {
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse and execute the command
		command, args := ParseCommand(line)
		lastExitCode = builtin.HandleCommands(command, args)

		// Exit if the command was "exit"
		if command == "exit" {
			os.Exit(lastExitCode)
		}
	}
}
