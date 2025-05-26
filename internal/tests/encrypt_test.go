package tests

import (
	"testing"

	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal"
	"github.com/stretchr/testify/assert"
)

func Test_CryptCommand_Execute(t *testing.T) {

	// Arrange
	sut := internal.CreateEncryptMessageCommand()
	sut.SetArgs([]string{"--message", "Hello World"})
	// Act
	err := sut.Execute()

	// Assert
	assert.NoError(t, err)
}
