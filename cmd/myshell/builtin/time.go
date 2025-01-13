package builtin

import (
	"fmt"
	"time"
)

var yellow = "\x1b[38;2;255;255;0m"
var defaultColor = "\x1b[0m"

func Time(args []string) {
	fmt.Println(yellow + "Time started" + defaultColor)
	start := time.Now()
	HandleCommands(args[0], args[1:])
	elapsed := time.Since(start)
	fmt.Println(yellow + "Elapsed time: " + elapsed.String() + defaultColor)
}
