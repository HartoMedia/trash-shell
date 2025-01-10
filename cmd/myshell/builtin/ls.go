package builtin

import (
	"fmt"
	"os"
)

func Ls(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Print("ls: error reading directory")
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
