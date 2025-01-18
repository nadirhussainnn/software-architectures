package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	conn, err := rpc.Dial("tcp", ":8080") // Note the rpc.Dial, in general we do net.Listen

	if err != nil {
		log.Print("Error occured while connecting", err)
	}
	defer conn.Close()

	// Calling the method
	var reply string
	err = conn.Call("Calendar.Today", "", &reply)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return
	}

	// Print the response
	fmt.Println("Server Response:", reply)

	// WHEN WE do net.Listen
	// response_reader := bufio.NewReader(conn)
	// for {
	// 	r, err := response_reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Print("Error occured while reading response", err)
	// 		continue
	// 	}
	// 	log.Print(r)
	// }
}
