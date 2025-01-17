package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Print("Connection could not be made")
	}

	defer conn.Close()

	log.Print("Connected to server. Type your message")
	// read from console
	reader := bufio.NewReader(os.Stdin)

	// read from connection
	server_response := bufio.NewReader(conn)
	for {
		log.Print(">>")

		input, err := reader.ReadString('\n')
		if err != nil {

		}

		// send input of keyboard to server
		_, err = conn.Write([]byte(input))

		if err != nil {
			log.Print("Could not write to server channel")
		}

		response, err := server_response.ReadString('\n')
		if err != nil {
			log.Print("Server did not responded!", err)
		}
		log.Print("Server says: ", response)

	}
}
