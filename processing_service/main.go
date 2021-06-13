package main

import (
	"fmt"
	"mateo/service/services/nats"
)

func main() {
	fmt.Println("Starting Microservice")

	for {
		nats.ListenToMessages()
	}
}
