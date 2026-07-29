package internal

// cryptCommand A base command type for crypt commands.
type cryptCommand struct {
	Passphrase string `cobra-x:"-p|--passphrase, help='The passphrase protecting a message.'"`
}
