package fileio

import (
	"fmt"
	"io/ioutil"
	"time"
)

func SaveFile(originalFileName, path string, file []byte) {
	savePath := fmt.Sprintf("%s/%s-%d", path, originalFileName, time.Now().Unix())
	ioutil.WriteFile(savePath, file, 0644)
}
