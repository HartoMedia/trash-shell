package builtin

import (
	"fmt"
	"os"
)

func Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}
