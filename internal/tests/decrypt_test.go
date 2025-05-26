package tests

import (
	"testing"

	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal"
	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_DecryptCommand_Execute(t *testing.T) {

	// Arrange
	const expectedMessage = "Hello World"

	encryptedMessage, _ := utils.CaptureStdout(t, "", func() error {
		encryptCommand := internal.CreateEncryptMessageCommand()
		encryptCommand.SetArgs([]string{"--message", expectedMessage})
		return encryptCommand.Execute()
	})

	decryptedMessage, err := utils.CaptureStdout(t, encryptedMessage, func() error {
		decryptCommand := internal.CreateDecryptMessageCommand()
		return decryptCommand.Execute()
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, decryptedMessage)
}
