package builtin

import (
	"fmt"
	"os"
)

func Mv(args []string) {
	if len(args) < 2 {
		return
	}
	err := os.Rename(args[0], args[1])
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "mv: %s: %v\n", args[0], err)
		if err != nil {
			return
		}
	}
}
