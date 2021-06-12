package tests

import (
	"fmt"
	"io/ioutil"
	"mateo/service/services/processing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxExtraction(t *testing.T) {
	data, err := ioutil.ReadFile("/home/mateo/uniqcast_exercise/video.mp4")

	if err != nil {
		panic(err)
	}

	boxes := processing.GetFileBoxes(data)

	for index, data := range boxes {
		fmt.Print(fmt.Sprintf("Index: %d, BoxType: %s", index, data.BoxType.Name()))
	}

	//assert that we have 6 boxes in our mp4 file
	assert.Len(t, boxes, 6)

	println("Done.")
}
