package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("publisher")
	opts.SetWill("status/clients", "publisher disconnected unexpectedly", 1, true)
	conn := mqtt.NewClient(opts)

	if token := conn.Connect(); token.Wait() && token.Error() != nil {
		log.Print("Error occured while adding publisher to broker")
	}

	for i := 0; i < 5; i++ {
		payload := fmt.Sprintf("%.2f", rand.Float32()*50)

		token := conn.Publish("temperature", 1, true, payload) // topic, qos level, retained message bool, payload
		// I have set retained message to true so we make sure guaranteed delivery in case of offline
		// Only last message will be delivered. For all messages, we need to configure broker
		token.Wait()

		time.Sleep(time.Second * 3) // send after every 2 seconds
	}

	log.Print("It is non-blocking")
}
