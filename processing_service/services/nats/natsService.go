package nats

import (
	"mateo/service/services/fileio"
	"mateo/service/services/processing"

	"github.com/nats-io/nats.go"
)

const CONN_STRING = "localhost:8222"
const OUTPUT_PATH = "/home/mateo/uniqcast_exercise/out"
const CHANNEL = "fileChannel"

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
	nc, err := nats.Connect(CONN_STRING)

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
		//get the init segment in byte format
		bytes, err := processing.GetInitSegmentAsBytes(im.FileData)

		outMsg := OutgoingMessage{}

		//check if there were issues with getting the init segment
		if err != nil {
			//if there were errors, respond
			outMsg.Error = err.Error()
			outMsg.Success = false
			respondToMessage(connection, outMsg)
			return
		}

		//check if there were errors with saving
		path, err := fileio.SaveFile(im.FileName, OUTPUT_PATH, bytes)
		if err != nil {
			//if there were errors, respond
			outMsg.Error = err.Error()
			outMsg.Success = false
			respondToMessage(connection, outMsg)
			return
		}

		outMsg.Success = true
		outMsg.Path = path
		respondToMessage(connection, outMsg)

	})
}

func respondToMessage(connection *nats.EncodedConn, outMsg OutgoingMessage) {
	connection.Publish(CHANNEL, outMsg)
}
