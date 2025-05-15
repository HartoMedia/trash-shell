# TRASH


The *Totally Random and Aimless Shell* or for short ***TRASH*** is a simple shell written in Go. It is a POSIX-compliant implementation of a Unix shell that supports basic shell commands, environment variables, script execution, and has many built-in commands.

## Project Structure

The project is structured as follows:
- `cmd/myshell/main.go`: The main shell entry point
- `cmd/myshell/parser.go`: Command parsing logic
- `cmd/myshell/ps1.go`: Prompt handling
- `cmd/myshell/builtin/`: Directory containing all built-in commands
  - `builtins.go`: Registry of all built-in commands
  - `handleCommands.go`: Command execution logic
  - Individual files for each built-in command (cd.go, echo.go, etc.)


## Changelog
Check the [CHANGELOG](CHANGELOG.md) for the latest changes.

## Getting Started

### Prerequisites

- Go 1.22 or later

### Installation
1. Install Go:
   - Arch Linux: `sudo pacman -S go`
   - Ubuntu: `sudo apt install golang-go`
   - Fedora: `sudo dnf install golang`
   - MacOS: `brew install go`
   - Windows: `scoop install go`
   - Windows (manual): Download the binary from the [official website](https://golang.org/dl/).
2. Clone the repository:

3. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Program

To run the shell program, execute:
```sh
go run main.go
```

### Building the Program

To build the shell program, execute:
```sh
go build -o shell.exe
```

This will create an executable file named `shell` in the project directory.

### Running the Executable

To run the executable on Unix or Unix-like systems:
```sh
./shell.exe
```

To run the executable on Windows:
```sh
.\shell.exe
```

### Cleaning Up

To clean up the project directory, execute:
```sh
go clean
```

## License

This project is licensed under the GNU General Public License v3.0 License. See the [LICENSE](LICENSE) file for details.

## Builtins

### Standard Builtins
- `cd`: Change the current directory (POSIX compliant, supports `~`, `-`, and no arguments).
- `exit`: Exit the shell with optional exit code.
- `pwd`: Print the current working directory.
- `echo`: Display a line of text (POSIX compliant, supports `-n`, `-e`, and `-E` options).
- `type`: Display information about command type.
- `ls`: List directory contents.
- `mkdir`: Create directories.
- `rm`: Remove files or directories.
- `mv`: Move or rename files.
- `touch`: Change file timestamps or create empty files.

### POSIX Required Builtins
- `:`: Null command (does nothing, always succeeds).
- `.`: Source/execute commands from a file in the current shell.
- `export`: Set environment variables for child processes.
- `set`: Set or unset shell options and positional parameters.
- `unset`: Unset environment variables.
- `test` or `[`: Evaluate conditional expressions.

### Additional Builtins
- `thx`: Displays the Credits.
- `cls`/`clear`: Clear the screen.
- `dirs`: Display directory stack.
- `help`: Display help information.
- `sleep`: Pause for a specified amount of time.
- `time`: Display the time.

The shell also runs any other command found in the "PATH" environment variable.

## Using as Default Shell on Linux

To use TRASH as your default shell on Linux:

1. Build the shell and install it to a standard location:
   ```sh
   go build -o myshell cmd/myshell/main.go
   sudo cp myshell /usr/local/bin/
   ```

2. Add the shell to the list of allowed shells:
   ```sh
   echo "/usr/local/bin/myshell" | sudo tee -a /etc/shells
   ```

3. Change your default shell:
   ```sh
   chsh -s /usr/local/bin/myshell
   ```

4. Log out and log back in to use your new shell.

### Shell Scripts

TRASH supports shell scripts with shebang lines. Create a script with:

```sh
#!/usr/local/bin/myshell
echo "Hello from TRASH shell script!"
```

Make it executable and run it:
```sh
chmod +x script.sh
./script.sh
```

## Authors

Developer & Maintainer:
- [Harto](https://blog.harto.dev)

Designer:
- [LordBuilder](https://www.curseforge.com/members/lordbuilder/projects)
