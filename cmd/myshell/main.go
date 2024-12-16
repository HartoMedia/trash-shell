package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var builtins = []string{"exit", "echo", "type", "pwd", "cd"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			os.Exit(1)
		}

		s = strings.Trim(s, "\r\n")

		var commands []string
		command, argstr, _ := strings.Cut(s, " ")
		if strings.Contains(s, "\"") {
			re := regexp.MustCompile("\"(.*?)\"")
			commands = re.FindAllString(s, -1)
			for i := range commands {
				commands[i] = strings.Trim(commands[i], "\"")
			}
		} else if strings.Contains(s, "'") {
			re := regexp.MustCompile("'(.*?)'")
			commands = re.FindAllString(s, -1)
			for i := range commands {
				commands[i] = strings.Trim(commands[i], "'")
			}
		} else {
			if strings.Contains(argstr, "\\") {
				re := regexp.MustCompile(`[^\\]+`)
				commands = re.Split(argstr, -1)
				for i := range commands {
					commands[i] = strings.ReplaceAll(commands[i], "\\", "")
				}
			} else {
				commands = strings.Fields(argstr)
			}

		}

		//for {
		//	start := strings.IndexAny(s, "'\"")
		//	if start == -1 {
		//		commands = append(commands, strings.Fields(s)...)
		//		break
		//	}
		//	ch := s[start]
		//	commands = append(commands, strings.Fields(s[:start])...)
		//	s = s[start+1:]
		//	end := strings.IndexByte(s, ch)
		//	token := s[:end]
		//	commands = append(commands, token)
		//	s = s[end+1:]
		//}

		switch command {

		case "exit":
			switch commands[1] {
			case "0":
				os.Exit(0)
			default:
				os.Exit(0)
			}

		case "echo":
			for i := 0; i < len(commands); i++ {
				fmt.Println(commands[i])
			}

		case "type":
			builtin_type(commands)

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)

		case "cd":
			switch commands[1] {
			case "~":
				os.Chdir(os.Getenv("HOME"))
			default:
				err := os.Chdir(commands[1])
				if err != nil {
					fmt.Printf("cd: %s: No such file or directory", commands[1])
					fmt.Println("")
				}
			}

		default:
			command := exec.Command(commands[0], commands[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", commands[0])
			}
		}
	}
}

func builtin_type(commands []string) {
	if len(commands) < 2 {
		return
	} else {
		for _, b := range builtins {
			if b == commands[1] {
				fmt.Println(commands[1] + " is a shell builtin")
				return
			}

		}
		env := os.Getenv("PATH")
		path := strings.Split(env, ":")
		for _, p := range path {
			if _, err := os.Stat(p + "/" + commands[1]); err == nil {
				fmt.Println(commands[1] + " is " + p + "/" + commands[1])
				return
			}
		}
		fmt.Println(commands[1] + ": not found")
	}
}
