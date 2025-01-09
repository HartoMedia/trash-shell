package builtin

import (
	"fmt"
	"os"
	"strconv"
)

func Exit(args []string) {
	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		os.Exit(n)
	} else {
		os.Exit(0)
	}
}
