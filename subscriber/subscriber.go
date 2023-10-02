package main

import (
	"fmt"
	"syscall"

	//import "log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang" // importo la libreria paho per gestire MQTT da Go
)

func onReceiveMessage(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Topic: %v\tPayload: %s\n", message.Topic(), message.Payload())
}

func main() {
	// definisco delle costanti per la connessione al broker MQTT
	const (
		broker = "127.0.0.1"
		port   = 1883
		hostId = "subscriber" // ID del client
		topic  = "test"
	)

	opts := mqtt.NewClientOptions()                          // tramite questa variabile verranno impostate tutte le opzioni utili per la connessione al broker MQTT
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port)) // imposto l'indirizzo del broker
	opts.SetClientID(hostId)                                 // imposto l'ID del client
	//opts.SetUsername("username")                           // imposto lo username (se richiesto)
	//opts.SetPassword("password")                           // imposto la password (se richiesta)
	client := mqtt.NewClient(opts) // inizializzo il client

	connectionToken := client.Connect() // effettuo la connessione

	if connectionToken.Wait() && connectionToken.Error() != nil {
		fmt.Println(connectionToken.Error()) // si può provare la generazione di questo messaggio d'errore ad esempio indicando un IP o una porta sbagliati
	} else {
		fmt.Println("Connected to the MQTT broker!")
	}

	subscriptionToken := client.Subscribe(topic, 0, onReceiveMessage) // faccio la subscription al topic; il secondo parametro è il QoS mentre il terzo è la funzione da chiamare quando si riceve un messaggio

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
