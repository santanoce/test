package libmqtt

import (
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Broker = "127.0.0.1"
const Port = 1883

func Connect(clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions()                          // tramite questa variabile verranno impostate tutte le opzioni utili per la connessione al broker MQTT
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", Broker, Port)) // imposto l'indirizzo del broker
	opts.SetClientID(clientID)                               // imposto l'ID del client
	//opts.SetUsername("username")                           // imposto lo username (se richiesto)
	//opts.SetPassword("password")                           // imposto la password (se richiesta)

	client := mqtt.NewClient(opts) // inizializzo il client

	connectionToken := client.Connect() // effettuo la connessione

	if connectionToken.Wait() && connectionToken.Error() != nil {
		fmt.Println(connectionToken.Error()) // si può provare la generazione di questo messaggio d'errore ad esempio indicando un IP o una porta sbagliati
	} else {
		fmt.Println("Connected to the MQTT broker!")
	}

	return client
}

func PublishRandomTemp(client mqtt.Client, topic string) {
	for {
		tmp := rand.Intn(40)
		text := fmt.Sprintf("%d", tmp)
		token := client.Publish(topic, 0, true, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
