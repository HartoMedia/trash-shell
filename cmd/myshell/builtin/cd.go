package builtin

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: missing argument")
		return
	}
	if args[0] == "~" {
		err := os.Chdir(os.Getenv("HOME"))
		if err != nil {
			return
		}
	} else if err := os.Chdir(args[0]); os.IsNotExist(err) {
		fmt.Println("cd: " + args[0] + ": No such file or directory")
	} else if err != nil {
		fmt.Println(err)
	}
}
