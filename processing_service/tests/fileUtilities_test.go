package tests

import (
	"fmt"
	"mateo/service/services/fileio"
	"testing"
)

func TestFileSave(t *testing.T) {
	file := []byte("test file")
	savePath := "../../"
	originalName := "testFile"

	fileio.SaveFile(originalName, savePath, file)

	fmt.Println("done")
}
