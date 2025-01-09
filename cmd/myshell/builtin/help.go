package builtin

import "fmt"

func Help(args []string) {
	builtins := GetBuiltins()
	fmt.Print("Builtins: " + builtins[0])
	if len(args) > 0 {
		switch args[0] {
		case "exit":
			fmt.Println(" - Exit the shell")
		case "echo":
			fmt.Println(" - Print arguments to the shell")
		case "type":
			fmt.Println(" - Print the type of the command")
		case "cls":
			fmt.Println(" - Clear the screen")
		case "pwd":
			fmt.Println(" - Print the current working directory")
		case "cd":
			fmt.Println(" - Change the current working directory")
		case "thx":
			fmt.Println(" - Print a thank you message")
		case "dirs":
			fmt.Println(" - Print the current working directory")
		case "help":
			fmt.Println(" - Print this help message")
		default:
			fmt.Println(" - Builtin command not found")
		}
	}
}
