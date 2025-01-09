package builtin

import (
	"fmt"
	"os"
	"strings"
)

func Dirs() string {
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
