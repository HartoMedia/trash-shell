package builtin

import (
	"fmt"
	"strings"
)

// Echo implements the POSIX echo command
// It supports the -n option to suppress the trailing newline
// and escape sequences when -e is specified
func Echo(args []string) {
	// Process options
	suppressNewline := false
	enableEscapes := false
	processedArgs := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			// Process options
			for _, opt := range arg[1:] {
				switch opt {
				case 'n':
					suppressNewline = true
				case 'e':
					enableEscapes = true
				case 'E':
					enableEscapes = false
				}
			}
		} else {
			processedArgs = append(processedArgs, arg)
		}
	}

	// Join arguments
	output := strings.Join(processedArgs, " ")

	// Process escape sequences if enabled
	if enableEscapes {
		output = processEscapes(output)
	}

	// Print output
	if suppressNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}

// processEscapes handles escape sequences in the string
func processEscapes(s string) string {
	var result strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'a': // Alert (bell)
				result.WriteByte('\a')
			case 'b': // Backspace
				result.WriteByte('\b')
			case 'c': // Suppress trailing newline
				return result.String()
			case 'e', 'E': // Escape
				result.WriteByte('\033')
			case 'f': // Form feed
				result.WriteByte('\f')
			case 'n': // New line
				result.WriteByte('\n')
			case 'r': // Carriage return
				result.WriteByte('\r')
			case 't': // Horizontal tab
				result.WriteByte('\t')
			case 'v': // Vertical tab
				result.WriteByte('\v')
			case '\\': // Backslash
				result.WriteByte('\\')
			default: // Just output the character
				result.WriteByte(s[i+1])
			}
			i += 2
		} else {
			result.WriteByte(s[i])
			i++
		}
	}
	return result.String()
}
