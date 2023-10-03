package main

import (
	"fmt"
	"libmqtt"
)

func main() {
	clientID := "publisher1"
	client := libmqtt.Connect(clientID)
	topic := fmt.Sprintf("temperature/%s", clientID)
	libmqtt.PublishRandomTemp(client, topic)
}
