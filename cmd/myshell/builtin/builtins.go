package builtin

var builtins = []string{"exit", "echo", "type", "cls", "pwd", "cd", "thx", "dirs", "help", "sleep"}

func GetBuiltins() []string {
	return builtins
}

func HandleBuiltins(command string, args []string) bool {
	is_builtin := true
	switch command {
	case "exit":
		Exit(args)
	case "echo":
		Echo(args)
	case "type":
		Type(args)
	case "cls":
		Cls()
	case "pwd":
		Pwd()
	case "cd":
		Cd(args)
	case "thx":
		Thx()
	case "dirs":
		Dirs()
	case "help":
		Help(args)
	case "sleep":
		Sleep(args)
	default:
		is_builtin = false
	}
	return is_builtin
}
