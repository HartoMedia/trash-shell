package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Prompt String 1

var blue = "\x1b[38;2;18;184;217m"
var magenta = "\x1b[38;2;187;44;135m"
var gitColor = "\x1b[38;2;240;80;51m"
var defaultColor = "\x1b[0m"

func Ps1() {
	_, err := fmt.Fprintf(os.Stdout, "%s%s %s %s %s%s%sΣ ", magenta, getUserName(), blue, getWorkDir(), gitColor, getGitBranch(), defaultColor)
	if err != nil {
		return
	}
}

func getWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	home := os.Getenv("HOME")
	if strings.HasPrefix(dir, home) {
		dir = "~" + strings.TrimPrefix(dir, home)
	}
	return fmt.Sprintf(strings.ReplaceAll(dir, "\\", "/"))
}

func getUserName() string {
	username := os.Getenv("USER")

	hostname := os.Getenv("HOST")

	return fmt.Sprintf("%s@%s", username, hostname)
}

func getGitBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	err := error(nil)
	cmd.Dir, err = os.Getwd()
	if err != nil {
		return ""
	}
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return ""
	}
	branch := strings.TrimSpace(string(out))
	return fmt.Sprintf("(%s) ", branch)
}
