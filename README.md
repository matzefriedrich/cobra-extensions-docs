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