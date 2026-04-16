# Go-Shell

A POSIX-compliant shell implementation written in Go. This project covers the core internals of a command-line interface, from process management and terminal I/O to complex string parsing and stream redirection.

---

## Features

### Core Mechanics
* **REPL & Prompt:** A standard Read-Eval-Print Loop that handles input and invalid command states.
* **Built-in Commands:** Native implementation of `echo`, `exit`, `pwd`, `type`, and `cd`.
* **Path Resolution:** Dynamically locates and executes external programs by searching the system `$PATH`.
* **Directory Navigation:** Full support for `cd` using absolute paths, relative paths, and home (`~`) directory shortcuts.

### Advanced Parsing
* **Quoting & Escaping:** Robust handling of single quotes (`'`), double quotes (`"`), and backslashes (`\`) across different contexts.
* **Quoted Executables:** Supports running programs located in paths that require quoting (e.g., paths with spaces).

### I/O & Control Flow
* **Redirection:** * Standard output and error redirection (`>`, `1>`, `2>`).
    * Appending support for both stdout and stderr (`>>`, `2>>`).
* **Pipelines:** Implementation of dual-command pipelines (`|`) for process chaining.

### Command Completion
* **Tab-Completion:** Intelligent completion for:
    * Built-in commands and their arguments.
    * Executables found in the system path.
    * Partial matches and multiple match selection logic.

---

## Usage

### Prerequisites
* Go 1.25 or higher

### Local Setup
To run the shell directly:
```sh
./your_program.sh
```

To build the binary:
```sh
go build -o goshell app/main.go
./goshell
```

---

## Roadmap

The primary architecture and POSIX-compliant features are complete. Future iterations will include:
* **Persistence:** Saving and navigating command history across sessions.
* **File Completion:** Expanding tab-completion to include local file paths and directories.

---

## Technical Learnings
This project involved a deep dive into Unix system calls, managing file descriptors for redirections, and handling process synchronization for pipelines in Go. I intentionally minimized the use of AI during development to ensure a grounded, first-principles understanding of low-level systems programming while reading documentation. 

### Acknowledgements
Special thanks to **CodeCrafters** for the structured challenge and excellent test suite (while it was free) that made building this shell a rewarding experience.
