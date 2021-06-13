package tests

import (
	"fmt"
	"mateo/service/services/fileio"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSave(t *testing.T) {
	file := []byte("test file")
	savePath := "../../"
	originalName := "testFile"

	result, err := fileio.SaveFile(originalName, savePath, file)

	assert.True(t, err == nil)
	assert.False(t, result == "")

	fmt.Println("done")
}
