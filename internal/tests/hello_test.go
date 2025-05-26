package tests

import (
	"testing"

	"github.com/matzefriedrich/cobra-extensions-docs/examples/internal"
	"github.com/stretchr/testify/assert"
)

func Test_HelloCommand_Execute(t *testing.T) {
	// Arrange
	sut := internal.CreateHelloCommand()
	sut.SetArgs([]string{"John Doe"})

	// Act
	err := sut.Execute()

	// Assert
	assert.NoError(t, err)
}
