![License](https://img.shields.io/github/license/matzefriedrich/cobra-extensions-docs)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/matzefriedrich/cobra-extensions-docs)

# cobra-extensions-docs

This repository contains example applications showcasing how to use the [`cobra-extensions`](https://github.com/matzefriedrich/cobra-extensions) package.

The examples serve as a reference for developers integrating `cobra-extensions` into their own CLI tools. Each example demonstrates practical usage of the features and utilities provided by the extensions package.

## Prerequisites

* Go 1.26+

## Getting Started

Clone the repository and navigate into the project folder:

```bash
git clone https://github.com/matzefriedrich/cobra-extensions-docs.git
cd cobra-extensions-docs
```

## Walk-through: Using cobra-extensions

The `cobra-extensions` package provides a declarative way to define CLI commands using Go structs and tags. This section provides a brief walk-through of the core concepts based on the examples in this repository.

### 1. Defining a Command

A command is represented by a struct that embeds `types.BaseCommand`. Use the `cobra-x` tag on this field to define the command's primary metadata, such as its name, a short help text, and an optional longer description.

```go
type helloCommand struct {
    types.BaseCommand `cobra-x:"hello, help='Prints a greeting to the specified name.'"`
}
```

### 2. Flags and Arguments

Flags and positional arguments are mapped to struct fields using the `cobra-x` tag.

*   **Flags**: Use the `--flag-name` (and optionally `-f` for shorthand) syntax in the tag.
*   **Positional Arguments**: Define a field (e.g., `Arguments`) that embeds `types.CommandArgs`. Subsequent fields in that struct are treated as positional arguments.

```go
type helloCommand struct {
    types.BaseCommand `cobra-x:"hello, help='Prints a greeting'"`
    
    // A string flag: --name or -n
    Name string `cobra-x:"--name, -n, help='The name to greet'"`

    // Positional arguments
    Arguments struct {
        types.CommandArgs
        Target string // This will be the first positional argument
    }
}
```

### 3. Implementing Execution Logic

To make the command functional, implement the `Execute` method. The `cobra-extensions` package automatically populates the struct fields from the command-line input before calling this method.

```go
func (c *helloCommand) Execute(ctx context.Context) {
    fmt.Printf("Hello %s %s!\n", c.Name, c.Arguments.Target)
}
```

### 4. Creating and Registering Commands

Use `commands.CreateTypedCommand` to transform your struct into a standard `*cobra.Command`.

```go
func CreateHelloCommand() *cobra.Command {
    instance := &helloCommand{
        Arguments: struct {
            types.CommandArgs
            Target string
        }{
            CommandArgs: types.NewCommandArgs(types.MinimumArgumentsRequired(1)),
        },
    }
    return commands.CreateTypedCommand(instance)
}
```

### 5. Command Groups and Sub-commands

You can nest commands to create complex CLI structures. The `AddGroupCommand` method in the `charmer` package allows you to group related commands together.

```go
charmer.NewCommandLineApplication("my-app", "A brief description").
    AddGroupCommand(CreateCryptCommand(), func(crypto types.CommandSetup) {
        crypto.AddCommand(
            CreateEncryptMessageCommand(),
            CreateDecryptMessageCommand())
    }).
    Execute(ctx)
```

### 6. Bootstrapping the Application

The `charmer` package provides a fluent API to set up and run your application, integrating your typed commands seamlessly.

```go
func main() {
    ctx := context.Background()
    
    err := charmer.NewCommandLineApplication("my-app", "A brief description").
        AddCommand(CreateHelloCommand()).
        Execute(ctx)
        
    if err != nil {
        log.Fatal(err)
    }
}
```

### 7. Built-in Utilities

The package includes several built-in utilities, such as a command to generate Markdown documentation for your CLI:

```go
app.AddCommand(commands.NewMarkdownDocsCommand(app))
```

## Build and Development

This project uses a `Makefile` to automate common tasks.

### Available Commands

List all available commands:
```bash
make help
```

| Command | Description |
| --- | --- |
| `make build` | Build all sample apps into the `build/` folder. |
| `make install-tools` | Install required tools (e.g., `golangci-lint`) to a local `bin/` folder. |
| `make lint-fix` | Run `golangci-lint` with the `--fix` flag. |
| `make clean` | Remove the `bin/` and `build/` folders. |

## Run examples

The sample applications are located in the `cmd/` directory. After building, you can find the binaries in the `build/` folder.

### Charmer Example
```bash
./build/charmer --help
```

### Simple Example
```bash
./build/simple --help
```

Alternatively, you can run them directly using `go run`:
```bash
go run cmd/charmer/main.go --help
```

--- 

Copyright 2025 - 2026 by Matthias Friedrich
