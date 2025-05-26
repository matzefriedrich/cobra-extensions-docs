package main

import (
	"os"

	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal"
	"github.com/matzefriedrich/cobra-extensions/pkg/charmer"
	"github.com/matzefriedrich/cobra-extensions/pkg/commands"

	"github.com/spf13/cobra"
)

func main() {

	app := charmer.NewRootCommand("simple-example", "")

	app.AddCommand(commands.NewMarkdownDocsCommand(app))
	app.AddCommand(internal.CreateHelloCommand())

	AddCryptCommands(app)

	err := app.Execute()
	if err != nil {
		return
	}

	os.Exit(0)
}

func AddCryptCommands(app *cobra.Command) {

	crypt := internal.CreateCryptCommand()

	crypt.AddCommand(
		internal.CreateEncryptMessageCommand(),
		internal.CreateDecryptMessageCommand(),
	)

	app.AddCommand(crypt)
}
