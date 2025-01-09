# TRASH


The *Totally Random and Aimless Shell* or for short ***TRASH*** is a simple shell written in Go. It is a minimalistic implementation of a Unix shell that supports basic shell commands and has some builtins.

## Project Structure

The project is structured as follows:
- `main.go`: Right now every thing is in this file. This follows that it is a bit unreadable. I will fix that (maybe) in the future.


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
go build -o shell
```

This will create an executable file named `shell` in the project directory.

### Running the Executable

To run the executable on Unix or Unix-like systems:
```sh
./shell
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

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Builtins

- `cd`: Change the current directory.
- `exit`: Exit the shell.
- `pwd`: Print the current working directory.
- `echo`: Display a line of text.
- `type`: Display information about command type.
- `thx`: Displays the Credits.
- runs every other command found in the "PATH" environment variable.

## Authors

Developer & Maintainer:
- [Harto](https://blog.harto.dev)

Designer:
- [LordBuilder](https://www.curseforge.com/members/lordbuilder/projects)