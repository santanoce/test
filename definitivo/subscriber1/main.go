package main

import (
	"fmt"
	"libmqtt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang" // importo la libreria paho per gestire MQTT da Go
)

var temperature = make(map[int]int)

func findMin(mappa map[int]int) int {
	min := math.MaxInt
	for _, v := range mappa {
		if v < min {
			min = v
		}
	}
	return min
}

func findMax(mappa map[int]int) int {
	max := math.MinInt
	for _, v := range mappa {
		if v > max {
			max = v
		}
	}
	return max
}

func onReceiveMessage(client mqtt.Client, message mqtt.Message) {
	topic := message.Topic()
	payload := message.Payload()
	tmp, _ := strconv.Atoi(fmt.Sprintf("%s", payload))
	fmt.Printf("Nuovo messaggio! Topic: %v\tPayload: %s\n", topic, payload)

	switch topic {
	case "temperature/publisher1":
		temperature[1] = tmp
	case "temperature/publisher2":
		temperature[2] = tmp
	case "temperature/publisher3":
		temperature[3] = tmp
	}
	fmt.Println("Le temperature attuali sono:", temperature)
	fmt.Println("La temperatura minima è:", findMin(temperature))
	fmt.Println("La temperatura massima è:", findMax(temperature))
}

func main() {
	clientID := "subscriber1"
	client := libmqtt.Connect(clientID)

	temperature[1] = 0
	temperature[2] = 0
	temperature[3] = 0

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
