package main

import (
	"fmt"
	"libmqtt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang" // importo la libreria paho per gestire MQTT da Go
)

func onReceiveMessage(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Topic: %v\tPayload: %s\n", message.Topic(), message.Payload())
}

func main() {
	clientID := "subscriber1"
	client := libmqtt.Connect(clientID)

	subscriptionToken := client.Subscribe("temperature/#", 0, onReceiveMessage) // faccio la subscription al topic; il secondo parametro è il QoS mentre il terzo è la funzione da chiamare quando si riceve un messaggio

	if subscriptionToken.Wait() && subscriptionToken.Error() != nil {
		fmt.Println(subscriptionToken.Error())
	} else {
		fmt.Println("Successfully subscribed to the topic!")
	}

	// Dopo aver fatto l'iscrizione al topic vogliamo che il client rimanga in ascolto; con il seguente codice il programma verrà chiuso solo al ricevimento di Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	fmt.Println("Signal caught - exiting . . .")
	client.Disconnect(1000)
	fmt.Println("Shutdown completed!")
}
