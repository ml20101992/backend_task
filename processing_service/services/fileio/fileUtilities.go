package fileio

import (
	"fmt"
	"io/ioutil"
	"time"
)

type FileSaveStatus struct {
	Success bool
	Err error
	Path string
}

func SaveFile(originalFileName, path string, file []byte) FileSaveStatus{
	savePath := fmt.Sprintf("%s/%s-%d", path, originalFileName, time.Now().Unix())
	err := ioutil.WriteFile(savePath, file, 0644)

	//there was an error while saving the file
	if err != nil {
		return FileSaveStatus{Success: false, Err: err}
	} else {
		return FileSaveStatus{Success: true, Path: savePath}
	}

}
