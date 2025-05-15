package builtin

import (
	"fmt"
	"os"
	"strings"
)

var builtins = []string{"exit", "echo", "type", "cls", "clear", "pwd", "cd", "thx", "dirs", "help", "sleep", "mkdir", "ls", "rm", "mv", "time", "touch"}

// Additional POSIX required builtins that need to be implemented
var posixBuiltins = []string{":", ".", "break", "continue", "eval", "exec", "export", "readonly", "return", "set", "shift", "test", "unset"}

func GetBuiltins() []string {
	allBuiltins := append([]string{}, builtins...)
	allBuiltins = append(allBuiltins, posixBuiltins...)
	return allBuiltins
}

// HandleBuiltins processes builtin commands and returns the exit code and whether the command was a builtin
func HandleBuiltins(command string, args []string) (int, bool) {
	exitCode := 0
	isBuiltin := true

	switch command {
	case "exit":
		// Exit is handled specially since it terminates the process
		if len(args) > 0 {
			Exit(args)
		} else {
			Exit([]string{"0"})
		}
		exitCode = 0 // This line won't be reached for exit, but included for completeness
	case "echo":
		Echo(args)
	case "type":
		Type(args)
	case "cls":
		Cls()
	case "clear":
		Cls()
	case "pwd":
		Pwd()
	case "cd":
		exitCode = CdWithExitCode(args)
	case "thx":
		Thx()
	case "dirs":
		Dirs()
	case "help":
		Help(args)
	case "sleep":
		Sleep(args)
	case "mkdir":
		exitCode = handleMkdir(args)
	case "ls":
		Ls(args)
	case "rm":
		exitCode = handleRm(args)
	case "mv":
		exitCode = handleMv(args)
	case "time":
		Time(args)
	case "touch":
		exitCode = handleTouch(args)
	// POSIX required builtins
	case ":": // Null command
		// Do nothing, always succeeds
	case ".": // Source command
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, ".: filename argument required")
			exitCode = 1
		} else {
			exitCode = handleSource(args[0], args[1:])
		}
	case "export":
		exitCode = handleExport(args)
	case "set":
		exitCode = handleSet(args)
	case "unset":
		exitCode = handleUnset(args)
	case "test", "[":
		exitCode = handleTest(args)
	default:
		isBuiltin = false
	}

	return exitCode, isBuiltin
}

// Helper functions for built-ins that need to return exit codes

func handleMkdir(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "mkdir: missing operand")
		return 1
	}

	Mkdir(args)
	return 0 // Assume success for now
}

func handleRm(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "rm: missing operand")
		return 1
	}

	Rm(args)
	return 0 // Assume success for now
}

func handleMv(args []string) int {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "mv: missing file operand")
		return 1
	}

	Mv(args)
	return 0 // Assume success for now
}

func handleTouch(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "touch: missing file operand")
		return 1
	}

	Touch(args)
	return 0 // Assume success for now
}

func handleSource(filename string, args []string) int {
	// This is a placeholder for the source command
	fmt.Fprintf(os.Stderr, ".: %s: file not found\n", filename)
	return 1
}

func handleExport(args []string) int {
	if len(args) == 0 {
		// List all exported variables
		for _, env := range os.Environ() {
			fmt.Println("export " + env)
		}
		return 0
	}

	// Process each argument
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			// Set variable
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) == 2 {
				os.Setenv(parts[0], parts[1])
			}
		} else {
			// Export existing variable
			val := os.Getenv(arg)
			if val == "" {
				// Variable doesn't exist, but this is not an error in POSIX
				os.Setenv(arg, "")
			}
		}
	}

	return 0
}

func handleSet(args []string) int {
	// This is a placeholder for the set command
	// In a real implementation, this would set shell options
	return 0
}

func handleUnset(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "unset: not enough arguments")
		return 1
	}

	// Unset each variable
	for _, arg := range args {
		os.Unsetenv(arg)
	}

	return 0
}

func handleTest(args []string) int {
	// This is a placeholder for the test command
	// In a real implementation, this would evaluate expressions
	if len(args) == 0 {
		return 1 // Empty test is false
	}
	return 0 // Assume true for now
}
