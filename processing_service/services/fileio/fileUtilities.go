package fileio

import (
	"fmt"
	"io/ioutil"
	"time"
)

func SaveFile(originalFileName, path string, file []byte) (string, error) {
	savePath := fmt.Sprintf("%s/%s-%d", path, originalFileName, time.Now().Unix())
	err := ioutil.WriteFile(savePath, file, 0644)

	//there was an error while saving the file
	if err != nil {
		return "", err
	} else {
		return savePath, nil
	}

}
