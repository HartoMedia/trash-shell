package builtin

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// HandleCommands executes the given command with arguments and returns the exit code
func HandleCommands(command string, args []string) int {
	// Handle environment variable expansion in arguments
	expandedArgs := make([]string, len(args))
	for i, arg := range args {
		expandedArgs[i] = expandEnvVars(arg)
	}
	args = expandedArgs

	// Parse command and arguments
	exitCode, isBuiltin := HandleBuiltins(command, args)
	if !isBuiltin {
		// Check if command exists
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
			return 127 // Command not found exit code
		}

		// Execute external command
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		// Set up environment variables
		cmd.Env = os.Environ()

		// Run the command
		if err := cmd.Run(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				return exitErr.ExitCode()
			}
			fmt.Fprintf(os.Stderr, "%s: error executing command\n", command)
			return 1 // Generic error exit code
		}
		return 0 // Success exit code
	}
	return exitCode
}

// expandEnvVars expands environment variables in the given string
// Handles $VAR and ${VAR} syntax
func expandEnvVars(s string) string {
	// Simple implementation for now
	result := s

	// Replace ${VAR} style variables
	for {
		start := strings.Index(result, "${")
		if start == -1 {
			break
		}

		end := strings.Index(result[start:], "}")
		if end == -1 {
			break
		}
		end += start

		varName := result[start+2 : end]
		varValue := os.Getenv(varName)

		result = result[:start] + varValue + result[end+1:]
	}

	// Replace $VAR style variables
	for {
		dollarIndex := strings.Index(result, "$")
		if dollarIndex == -1 || dollarIndex == len(result)-1 {
			break
		}

		// Find the end of the variable name
		endIndex := dollarIndex + 1
		for endIndex < len(result) && (isAlphaNumeric(result[endIndex]) || result[endIndex] == '_') {
			endIndex++
		}

		if endIndex == dollarIndex+1 {
			// Just a $ followed by a non-alphanumeric character
			break
		}

		varName := result[dollarIndex+1 : endIndex]
		varValue := os.Getenv(varName)

		result = result[:dollarIndex] + varValue + result[endIndex:]
	}

	return result
}

// isAlphaNumeric returns true if the given byte is a letter, digit, or underscore
func isAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_'
}
