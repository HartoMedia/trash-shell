package builtin

import (
	"fmt"
	"os"
)

func Touch(args []string) {
	for index, filename := range args {
		file, err := os.Create(filename)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "touch: %s: %v\n", filename, "at index", index, err)
			if err != nil {
				return
			}
		}
		err = file.Close()
		if err != nil {
			return
		}
	}
}
