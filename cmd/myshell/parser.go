package main

import (
	"regexp"
	"strings"
)

func parseCommand(input string) (string, []string) {
	// Handle quoted strings and escaped characters
	re := regexp.MustCompile(`\\.|"(.*?)"|'(.*?)'|\S+`)
	matches := re.FindAllString(input, -1)

	var args []string
	for _, match := range matches {
		if strings.HasPrefix(match, "\"") && strings.HasSuffix(match, "\"") {
			args = append(args, strings.Trim(match, "\""))
		} else if strings.HasPrefix(match, "'") && strings.HasSuffix(match, "'") {
			args = append(args, strings.Trim(match, "'"))
		} else {
			args = append(args, unescapeBackslashes(match))
		}
	}

	if len(args) == 0 {
		return "", nil
	}

	command := args[0]
	args = args[1:]
	return command, args
}

func unescapeBackslashes(input string) string {
	// Replace escaped characters (e.g., \\ -> \, \" -> ", etc.)
	return strings.ReplaceAll(input, "\\", "")
}
