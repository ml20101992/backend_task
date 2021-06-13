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

	result := fileio.SaveFile(originalName, savePath, file)

	assert.True(t, result.Success)

	fmt.Println("done")
}
