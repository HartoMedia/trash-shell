package builtin

import (
	"strconv"
	"time"
)

func Sleep(args []string) {
	if len(args) < 1 {
		println("sleep: missing operand")
		return
	}
	seconds, err := strconv.Atoi(args[0])
	if err != nil {
		println("sleep: invalid operand")
		return
	}

	time.Sleep(time.Duration(seconds) * time.Second)
}
