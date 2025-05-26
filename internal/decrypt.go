package internal

import (
	"io"
	"os"

	"github.com/matzefriedrich/cobra-extensions/pkg/commands"
	"github.com/matzefriedrich/cobra-extensions/pkg/types"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/spf13/cobra"
)

var _ types.TypedCommand = &decryptMessageCommand{}

type decryptMessageCommand struct {
	cryptCommand
	use types.CommandName `flag:"decrypt" short:"Decrypt a message." long:"Reads an armored message from stdin and decrypts it."`
}

func CreateDecryptMessageCommand() *cobra.Command {
	instance := &decryptMessageCommand{cryptCommand: cryptCommand{}}
	return commands.CreateTypedCommand(instance)
}

func (d *decryptMessageCommand) Execute() {
	armored, _ := ReadArmoredMessagedFromStdin()
	message, _ := crypto.NewPGPMessageFromArmored(armored)
	decrypted, _ := crypto.DecryptMessageWithPassword(message, []byte(d.Passphrase))
	_, _ = os.Stdout.WriteString(decrypted.GetString())
}

func ReadArmoredMessagedFromStdin() (string, error) {
	data, err := io.ReadAll(os.Stdin)
	if err == nil {
		return string(data), nil
	}
	return "", err
}
