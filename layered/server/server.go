package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func processRequest(request string) string {

	output := fmt.Sprintf("%v -- %v:%v", strings.ToUpper(request), time.Now().Hour(), time.Now().Minute())
	return output

}
func handleConnection(conn net.Conn) {

	defer conn.Close()

	log.Print("Client connected ", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			log.Print("Error occured in reading content", err)
		}
		request = strings.TrimSpace(request)
		response := processRequest(request)

		_, err = conn.Write([]byte(response + "\n"))

		if err != nil {
			log.Print("Error occured in writing response to connection")
			return
		}
	}
}
func main() {

	PORT := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))

	if err != nil {
		log.Print("Error occured")
	}

	log.Print(fmt.Sprintf("Server is listening at PORT %v", PORT))
	defer listener.Close()

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Print("Error occured in accepting connection")
		}

		go handleConnection(conn)
	}
}
