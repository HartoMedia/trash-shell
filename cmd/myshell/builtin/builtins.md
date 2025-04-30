# Core Shell Management
- [x] **exit**: Exits the shell.  
  Example: `exit` or `exit <code>`.
- [x] **cd**: Changes the current directory.  
  Example: `cd <directory>` or `cd ~`.
- [x] **pwd**: Prints the current working directory.  
  Example: `pwd`.
- [ ] **help**: Displays help information for builtins or shell usage.  
  Example: `help` or `help <builtin>`.
- [x] **echo**: Prints text to the terminal.  
  Example: `echo "Hello, World!"`.
- [x] **clear**: Clears the hole stdOut.  
  Example: `clear`.

# Process Management
- [ ] **jobs**: Lists background jobs.  
  Example: `jobs`.
- [ ] **fg**: Brings a background job to the foreground.  
  Example: `fg <job_id>`.
- [ ] **bg**: Resumes a background job in the background.  
  Example: `bg <job_id>`.
- [ ] **kill**: Sends a signal to a process.  
  Example: `kill <pid>` or `kill -9 <pid>`.

# Shell Customization
- [ ] **alias**: Creates a shortcut for commands.  
  Example: `alias ll='ls -la'`.
- [ ] **unalias**: Removes an alias.  
  Example: `unalias ll`.
- [ ] **set**: Configures shell options or variables.  
  Example: `set -o vi` or `set VAR=value`.
- [ ] **unset**: Unsets a variable or option.  
  Example: `unset VAR`.

# File and Directory Management
- [x] **ls**: Lists files and directories.  
  Example: `ls` or `ls -la`.
- [x] **mkdir**: Creates a directory.  
  Example: `mkdir <directory>`.
- [ ] **rmdir**: Removes an empty directory.  
  Example: `rmdir <directory>`.
- [x] **rm**: Deletes files or directories.  
  Example: `rm <file>` or `rm -r <directory>`.
- [x] **mv**: Moves or renames files or directories.  
  Example: `mv <source> <destination>`.
- [ ] **cp**: Copies files or directories.  
  Example: `cp <source> <destination>`.
- [x] **touch**: Creates an empty file or updates the timestamp.  
  Example: `touch file.txt`.
- [ ] **cat**: Displays the contents of a file.  
  Example: `cat file.txt`.
- [ ] **head**: Displays the first few lines of a file.  
  Example: `head -n 5 file.txt`.
- [ ] **tail**: Displays the last few lines of a file.  
  Example: `tail -n 5 file.txt`.
- [ ] **find**: Searches for files and directories.  
  Example: `find . -name "*.txt"`.
- [ ] **basename**: Extracts the filename from a path.  
  Example: `basename /path/to/file.txt`.
- [ ] **dirname**: Extracts the directory from a path.  
  Example: `dirname /path/to/file.txt`.

# Redirection and Piping
- [ ] **>**: Redirects output to a file.  
  Example: `echo "data" > file.txt`.
- [ ] **>>**: Appends output to a file.  
  Example: `echo "data" >> file.txt`.
- [ ] **|**: Pipes the output of one command to another.  
  Example: `ls | grep "pattern"`.

# Debugging and Info
- [ ] **env**: Displays environment variables.  
  Example: `env`.
- [ ] **export**: Sets environment variables.  
  Example: `export VAR=value`.
- [ ] **history**: Shows command history.  
  Example: `history`.
- [x] **type**: Displays the type of a command.  
  Example: `type <command>`.
- [ ] **true**: Always returns success.  
  Example: `true`.
- [ ] **false**: Always returns failure.  
  Example: `false`.
- [x] **sleep**: Pauses execution for a specified time.  
  Example: `sleep 5`.
- [ ] **wait**: Waits for background processes to finish.  
  Example: `wait <pid>`.
- [ ] **printenv**: Prints all environment variables.  
  Example: `printenv`.
- [ ] **readonly**: Marks a variable as read-only.  
  Example: `readonly VAR=value`.

# User and System Management
- [ ] **whoami**: Prints the current user.  
  Example: `whoami`.
- [ ] **id**: Displays user and group IDs.  
  Example: `id`.
- [ ] **uptime**: Shows system uptime.  
  Example: `uptime`.
- [ ] **hostname**: Displays or sets the hostname.  
  Example: `hostname` or `hostname new-host`.
- [ ] **who**: Lists logged-in users.  
  Example: `who`.
- [ ] **groups**: Lists groups for the current user.  
  Example: `groups`.
- [ ] **umask**: Sets default file permissions.  
  Example: `umask 022`.
- [ ] **chown**: Changes file ownership.  
  Example: `chown user:group file.txt`.
- [ ] **chmod**: Changes file permissions.  
  Example: `chmod 755 file.txt`.

# Advanced Shell Features
- [ ] **let**: Evaluates arithmetic expressions.  
  Example: `let "a = 5 + 3"`.
- [ ] **eval**: Executes arguments as a command.  
  Example: `eval echo \$HOME`.
- [ ] **command**: Executes a command, bypassing aliases.  
  Example: `command ls`.
- [ ] **builtin**: Executes a shell builtin, bypassing overrides.  
  Example: `builtin echo "This is a builtin"`.
- [ ] **declare**: Declares variables and their attributes.  
  Example: `declare -i num=5`.

# Networking
- [ ] **curl**: Fetches data from a URL.  
  Example: `curl <url>`.
- [ ] **wget**: Downloads files from a URL.  
  Example: `wget <url>`.
- [ ] **ping**: Sends ICMP echo requests.  
  Example: `ping 8.8.8.8`.
- [ ] **nslookup**: Queries DNS information.  
  Example: `nslookup example.com`.
- [ ] **traceroute**: Traces the route packets take.  
  Example: `traceroute example.com`.
- [ ] **netstat**: Displays network connections.  
  Example: `netstat -an`.

# Scripting Utilities
- [ ] **read**: Reads input from the user.  
  Example: `read name`.
- [ ] **shift**: Shifts positional parameters in a script.  
  Example: `shift`.
- [ ] **getopts**: Parses command-line options in scripts.  
  Example: `getopts "a:b:" opt`.
- [ ] **printf**: Formats and prints text.  
  Example: `printf "Hello %s\n" "World"`.
- [ ] **test / [ ]**: Evaluates conditions.  
  Example: `[ -f file.txt ]`.

# Version Control and Collaboration
- [ ] **git** (wrapper): Optionally provide shortcuts for Git commands.  
  Example: `git status`.
- [ ] **scp**: Copies files between systems using SSH.  
  Example: `scp file.txt user@host:/path`.
- [ ] **ssh**: Opens an SSH connection.  
  Example: `ssh user@host`.

# Custom Debug/Meta Commands
- [ ] **dump**: Dumps internal state of the shell for debugging.  
  Example: `dump`.
- [ ] **builtin-list**: Lists all available builtins in the shell.  
  Example: `builtin-list`.
- [ ] **trace**: Enables or disables command tracing.  
  Example: `trace on` or `trace off`.
- [ ] **reload**: Reloads shell configuration or scripts.  
  Example: `reload ~/.trashrc`.

[Link](https://chatgpt.com/c/678123e1-8bf4-8000-bf11-5b889a2e0262)
