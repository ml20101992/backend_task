package nats

import (
	"mateo/service/services/processing"

	"github.com/nats-io/nats.go"
)

var connString = "localhost:8222"

type IncomingMessage struct {
	FileName string
	FileData []byte
}

type OutgoingMessage struct {
	Success bool
	Error   string
	Path    string
}

//Function used to establish connection to NATS server and to set EncodedConn
func connectToNats() *nats.EncodedConn {
	nc, err := nats.Connect(connString)

	if err != nil {
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		panic(err)
	}

	return ec
}

//method used to listen for messages
func ListenToMessages() {
	connection := connectToNats()

	connection.Subscribe("fileChannel", func(im *IncomingMessage) {
		bytes, err := processing.GetInitSegmentAsBytes(im.FileData)

	})
}
