package builtin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Cd changes the current directory
// This is kept for backward compatibility
func Cd(args []string) {
	CdWithExitCode(args)
}

// CdWithExitCode changes the current directory and returns an exit code
// It follows POSIX specifications:
// - With no arguments, changes to $HOME
// - With "-" as argument, changes to $OLDPWD
// - Handles "~" and "~/" path expansion
// - Sets PWD and OLDPWD environment variables
func CdWithExitCode(args []string) int {
	// Get current directory for OLDPWD
	oldpwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cd: error getting current directory:", err)
		return 1
	}

	// Determine target directory
	var targetDir string

	if len(args) == 0 {
		// No args, go to $HOME
		targetDir = os.Getenv("HOME")
		if targetDir == "" {
			fmt.Fprintln(os.Stderr, "cd: HOME not set")
			return 1
		}
	} else if args[0] == "-" {
		// Go to previous directory
		targetDir = os.Getenv("OLDPWD")
		if targetDir == "" {
			fmt.Fprintln(os.Stderr, "cd: OLDPWD not set")
			return 1
		}
		// Print the directory when using cd -
		fmt.Println(targetDir)
	} else if args[0] == "~" || strings.HasPrefix(args[0], "~/") {
		// Handle home directory expansion
		home := os.Getenv("HOME")
		if home == "" {
			fmt.Fprintln(os.Stderr, "cd: HOME not set")
			return 1
		}

		if args[0] == "~" {
			targetDir = home
		} else {
			// Replace ~/ with $HOME/
			targetDir = filepath.Join(home, args[0][2:])
		}
	} else {
		// Regular directory
		targetDir = args[0]
	}

	// Change to the target directory
	if err := os.Chdir(targetDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", targetDir)
		} else {
			fmt.Fprintf(os.Stderr, "cd: %s: %v\n", targetDir, err)
		}
		return 1
	}

	// Get the new current directory
	newpwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cd: error getting new directory:", err)
		return 1
	}

	// Set PWD and OLDPWD environment variables
	os.Setenv("OLDPWD", oldpwd)
	os.Setenv("PWD", newpwd)

	return 0
}
