package builtin

import (
	"fmt"
	"os"
)

func Rm(args []string) {

	// check for flags like -r or -f and store them in a boolean each and also if the are toghether like -rf or -fr
	var r, f = false, false
	for _, arg := range args {
		if arg == "-r" {
			r = true
		} else if arg == "-f" {
			f = true
		}
	}

	// remove the flags from the args slice
	for i := 0; i < len(args); i++ {
		if args[i] == "-r" || args[i] == "-f" {
			args = append(args[:i], args[i+1:]...)
		}
	}

	// check if there are any arguments left
	if len(args) == 0 {
		fmt.Println("rm: missing operand")
		return
	}

	// get the path from the args
	path := args[0]

	//check if there thing to remove is a folder
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if file.IsDir() {
		if r {
			if f {
				err := os.RemoveAll(path)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("rm: cannot remove ", path, ": Is a directory")
		}
	} else {
		err := os.Remove(args[0])
		if err != nil {
			fmt.Print(err)
		}
	}
}
