package main

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("subscriber")
	opts.SetCleanSession(false) // For persistent session
	opts.SetKeepAlive(30)       // ping every 30 second to tell broker you are connected to it. If will message is configured, it will be triggered. Else broker knows this client is disconnected

	conn := mqtt.NewClient(opts)

	if token := conn.Connect(); token.Wait() && token.Error() != nil {
		log.Print("Error occured while registering subscriber to broker")
	}

	conn.Subscribe("temperature", 1, func(c mqtt.Client, m mqtt.Message) {

		log.Print(m.Topic())
		log.Print("Payload :", string(m.Payload()))
	})

	conn.Subscribe("status/clients", 1, func(c mqtt.Client, m mqtt.Message) {

		log.Print(m.Topic())
		log.Print("Payload :", string(m.Payload()))
	})
	// conn.Unsubscribe("temperature")				// to un-subscribe
	// conn.Disconnect(5000) 						// disconnect but wait 5000 millisecond
	select {} // wait infinite time for messages

}
