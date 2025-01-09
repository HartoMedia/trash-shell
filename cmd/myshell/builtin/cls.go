package builtin

import "fmt"

func Cls() {
	fmt.Print("\033[H\033[2J")
}
