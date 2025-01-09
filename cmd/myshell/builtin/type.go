package builtin

import (
	"fmt"
	"os"
	"strings"
)

var builtins = []string{"exit", "echo", "type", "pwd", "cd", "thx", "cls"}

func Type(commands []string) {
	if len(commands) == 0 {
		fmt.Println("type: missing argument")
		return
	}
	for _, b := range builtins {
		if b == commands[0] {
			fmt.Println(commands[0] + " is a shell builtin")
			return
		}
	}
	env := os.Getenv("PATH")
	path := strings.Split(env, ":")
	for _, p := range path {
		if _, err := os.Stat(p + "/" + commands[0]); err == nil {
			fmt.Println(commands[0] + " is " + p + "/" + commands[0])
			return
		}
	}
	fmt.Println(commands[0] + ": not found")
}
