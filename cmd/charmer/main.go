package main

import (
	"context"
	"log"

	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal"
	"github.com/matzefriedrich/cobra-extensions/pkg/charmer"
	"github.com/matzefriedrich/cobra-extensions/pkg/types"
)

func main() {

	ctx := context.Background()

	err :=
		charmer.NewCommandLineApplication("charmer-example", "A sample application to showcase the charmer package.").
			AddCommand(internal.CreateHelloCommand()).
			AddGroupCommand(internal.CreateCryptCommand(), func(crypto types.CommandSetup) {
				crypto.AddCommand(
					internal.CreateEncryptMessageCommand(),
					internal.CreateDecryptMessageCommand())
			}).
			Execute(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
