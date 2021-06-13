package nats

import (
	"fmt"
	"mateo/service/services/fileio"
	"mateo/service/services/processing"
	"sync"

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
	nc, err := nats.Connect(nats.DefaultURL)

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

	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := connection.Subscribe("fileChannel", func(im *IncomingMessage) {
		fmt.Println("Message Recieved...")
		//get the init segment in byte format
		bytes, err := processing.GetInitSegmentAsBytes(im.FileData)

		outMsg := OutgoingMessage{}

		//check if there were issues with getting the init segment
		if err != nil {
			//if there were errors, respond
			outMsg.Error = err.Error()
			outMsg.Success = false
			respondToMessage(connection, outMsg)
		} else {
			//check if there were errors with saving
			path, err := fileio.SaveFile(im.FileName, OUTPUT_PATH, bytes)
			if err != nil {
				//if there were errors, respond
				outMsg.Error = err.Error()
				outMsg.Success = false
				respondToMessage(connection, outMsg)
			} else {
				//everything is fine, respond to message
				outMsg.Success = true
				outMsg.Path = path
				respondToMessage(connection, outMsg)
			}
			connection.Close()
			wg.Done()
		}
	}); err != nil {
		panic(err)
	}

	fmt.Println("Waiting for messages")
	wg.Wait()
}

func respondToMessage(connection *nats.EncodedConn, outMsg OutgoingMessage) {
	connection.Publish(CHANNEL, outMsg)
}
