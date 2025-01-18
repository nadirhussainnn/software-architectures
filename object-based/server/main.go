package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type CalendarInterface interface {
	Today(args string, reply *string) error
}

type Calendar struct{}

func (c *Calendar) Today(args string, reply *string) error {
	*reply = time.Now().String()
	return nil
}

func main() {

	// register the object

	var calendar CalendarInterface = new(Calendar)

	err := rpc.Register(calendar)
	if err != nil {
		log.Print("Error registering service:", err)
		return
	}
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Print("Error occured while listening", err)
		return
	}
	log.Print("Server listening for object to be called")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		log.Print("Client connected ", conn.RemoteAddr())

		if err != nil {
			log.Print("Error occured while accepting connection")
			continue
		}

		// _, err = conn.Write([]byte("You are connected to me, and you can call my methods if you how i allow you!\n"))
		// if err != nil {
		// 	log.Print("Could not send response")
		// 	return
		// }
		go rpc.ServeConn(conn)
	}
}
