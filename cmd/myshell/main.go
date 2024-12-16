package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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

		var args []string
		command, argstr, _ := strings.Cut(s, " ")
		if strings.Contains(s, "\"") {
			re := regexp.MustCompile("\"(.*?)\"")
			args = re.FindAllString(s, -1)
			for i := range args {
				args[i] = strings.Trim(args[i], "\"")
			}
		} else if strings.Contains(s, "'") {
			re := regexp.MustCompile("'(.*?)'")
			args = re.FindAllString(s, -1)
			for i := range args {
				args[i] = strings.Trim(args[i], "'")
			}
		} else {
			if strings.Contains(argstr, "\\") {
				re := regexp.MustCompile(`[^\\]+`)
				args = re.Split(argstr, -1)
				for i := range args {
					args[i] = strings.ReplaceAll(args[i], "\\", "")
				}
			} else {
				args = strings.Fields(argstr)
			}
		}

		switch command {
		case "exit":
			n, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(n)

		case "echo":
			fmt.Println(strings.Join(args, " "))

		case "type":
			builtin_type(args)

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)

		case "cd":
			if args[0] == "~" {
				os.Chdir(os.Getenv("HOME"))
			}
			if err := os.Chdir(args[0]); os.IsNotExist(err) {
				fmt.Println(command + ": " + args[0] + ": No such file or directory")
				break
			} else if err != nil {
				fmt.Println(err)
			}

		default:
			_, err = exec.LookPath(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
				break
			}
			command := exec.Command(command, args...)
			command.Stdin = os.Stdin
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", args[0])
			}
		}
	}
}

func builtin_type(commands []string) {
	for _, b := range builtins {
		if b == commands[0] {
			fmt.Println(commands[0] + " is a shell builtin")
			return
		}

	}
	env := os.Getenv("PATH")
	path := strings.Split(env, ":")
	for _, p := range path {
		if _, err := os.Stat(p + "/" + commands[0]); err == nil {
			fmt.Println(commands[0] + " is " + p + "/" + commands[0])
			return
		}
	}
	fmt.Println(commands[0] + ": not found")
}
