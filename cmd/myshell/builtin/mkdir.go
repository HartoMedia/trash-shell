package builtin

import (
	"fmt"
	"os"
)

func Mkdir(args []string) {
	if len(args) == 0 {
		return
	}
	err := os.Mkdir(args[0], 0755)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "mkdir: %s: %v\n", args[0], err)
		if err != nil {
			return
		}
	}
}
