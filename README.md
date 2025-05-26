[![Go Report Card](https://goreportcard.com/badge/github.com/matzefriedrich/cobra-extensions-docs)](https://goreportcard.com/report/github.com/matzefriedrich/cobra-extensions-docs)
![License](https://img.shields.io/github/license/matzefriedrich/cobra-extensions-docs)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/matzefriedrich/cobra-extensions-docs)

# cobra-extensions-docs

This repository contains example applications showcasing how to use the [`cobra-extensions`](https://github.com/matzefriedrich/cobra-extensions) package.

The examples serve as a reference for developers integrating `cobra-extensions` into their own CLI tools. Each example demonstrates practical usage of the features and utilities provided by the extensions package.

## Prerequisites

* Go 1.24+

## Getting Started

Clone the repository and navigate into any example to explore its functionality:

```bash
git clone https://github.com/matzefriedrich/cobra-extensions-docs.git
cd cobra-extensions-docs
````

## Run examples

Each example has its own entry module. For instance, the following command builds and runs the `charmer` example app:

```sh
go run cmd/charmer/main.go --help
```