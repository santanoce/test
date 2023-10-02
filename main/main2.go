package main

import (
	"fmt"
	//import "log"

	mqtt "github.com/eclipse/paho.mqtt.golang" // importo la libreria paho per gestire MQTT da Go
)

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

	if !connectionToken.Wait() || connectionToken.Error() == nil {
		fmt.Println("Connected to the MQTT broker!")
	} else {
		fmt.Println(connectionToken.Error()) // si può provare la generazione di questo messaggio d'errore ad esempio indicando un IP o una porta sbagliati
	}

}
