package main

import (
	"fmt"
	"libmqtt"
)

func main() {
	clientID := "publisher3"
	client := libmqtt.Connect(clientID)
	topic := fmt.Sprintf("temperature/%s", clientID)
	libmqtt.PublishRandomTemp(client, topic)
}
