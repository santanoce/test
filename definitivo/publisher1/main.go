package main

import (
	"fmt"
	"libmqtt"
	"math/rand"
	"time"
	//mqtt "github.com/eclipse/paho.mqtt.golang" // importo la libreria paho per gestire MQTT da Go
)

func main() {
	clientID := "publisher1"
	client := libmqtt.Connect(clientID)

	for {
		tmp := rand.Intn(40)
		text := fmt.Sprintf("%d", tmp)
		token := client.Publish("temperature", 0, true, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
