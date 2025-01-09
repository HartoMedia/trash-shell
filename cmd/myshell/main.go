package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
)

var builtins = []string{"exit", "echo", "type", "pwd", "cd", "thx"}

var blau = "\x1b[38;2;18;184;217m"
var magenta = "\x1b[38;2;187;44;135m"
var gitColor = "\x1b[38;2;240;80;51m"
var defaultColor = "\x1b[0m"

func main() {
	for {
		fmt.Fprintf(os.Stdout, "%s%s %s %s %s%s%sΣ ", magenta, getUserName(), blau, getWorkDir(), gitColor, getGitBranch(), defaultColor)

		// Wait for user input
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			os.Exit(1)
		}

		s = strings.Trim(s, "\r\n")

		if len(s) == 0 {
			continue
		}

		// Parse command and arguments
		command, args := parseCommand(s)

		switch command {
		case "exit":
			if len(args) > 0 {
				n, err := strconv.Atoi(args[0])
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(n)
			} else {
				os.Exit(0)
			}

		case "echo":
			fmt.Println(strings.Join(args, " "))

		case "type":
			builtinType(args)

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)

		case "cd":
			if len(args) == 0 {
				fmt.Println("cd: missing argument")
				break
			}
			if args[0] == "~" {
				os.Chdir(os.Getenv("HOME"))
			} else if err := os.Chdir(args[0]); os.IsNotExist(err) {
				fmt.Println(command + ": " + args[0] + ": No such file or directory")
				break
			} else if err != nil {
				fmt.Println(err)
			}

		case "thx":
			fmt.Println("+--------------------------------------------------------------+\n" +
				"|                             THANK YOU                        |\n" +
				"+--------------------------------------------------------------+\n" +
				"| A big shoutout to *LordBuilder* for crafting the epic logo!  |\n" +
				"|                                                              |\n" +
				"|   Code for this shell was passionately written by *Harto*!   |\n" +
				"|                                                              |\n" +
				"| Thank you for using this shell. We hope you enjoy it as      |\n" +
				"| much as we enjoyed creating it.                              |\n" +
				"|                                                              |\n" +
				"| May your commands be swift, your scripts elegant, and your   |\n" +
				"| errors... minimal!                                           |\n" +
				"+--------------------------------------------------------------+\n")

		default:
			_, err = exec.LookPath(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
				break
			}
			cmd := exec.Command(command, args...)
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Printf("%s: error executing command\n", command)
			}
		}
	}
}

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

func builtinType(commands []string) {
	if len(commands) == 0 {
		fmt.Println("type: missing argument")
		return
	}
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

func getWorkDir() string {
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

func getUserName() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s@%s", strings.SplitN(u.Username, "\\", 2)[1], hostname)
}

func getGitBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	err := error(nil)
	cmd.Dir, err = os.Getwd()
	if err != nil {
		return ""
	}
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return ""
	}
	branch := strings.TrimSpace(string(out))
	return fmt.Sprintf("(%s) ", branch)
}
