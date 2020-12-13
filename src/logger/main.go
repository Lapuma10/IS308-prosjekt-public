package main

import (
	db "./db"
	//rabbit "./listener"
	rbMQ "./rbMQ"
	"log"
	"strings"
)

// NOTE FROM DANIEL: tried my best to update this legacy code with the new rbMQ code. 
func main(){
	// -- Setup for RBMQ communication
	// Get connection for this work session
	conn := rbMQ.GetConnection()
	defer conn.Close()

	ch := rbMQ.GetChannelFromConnection(conn)
	defer ch.Close()

	//rabbit.Listen()
	msgs := rbMQ.GetChannelDataByName(ch, "log")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			log.Printf("Logging message...")
			message := strings.Split(string(d.Body),":")
			db.InsertRecord(message[1],message[0])
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}