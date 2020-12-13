package rbMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)
/*
HOW TO USE THIS THING:
Parameters
1. a recieve-only channel initialized in main program 
2. the channel you get from rbMQ.GetChannelFromConnection(conn *amqp.Connection)

How to use
- Initialize a receive only channel in main program
- Get a connection to rbMQ
- From connection to rbMQ, get Channel
- Pass both into this function and run it as a go routine
- Send messages to the channel like this "in <- message"
- Remember to use a message formatter

What it does
- waits for input from the receive-only channel
- publishes message to the "log" queue
*/
func LogMessage(in <-chan string, ch *amqp.Channel) {
	for {
		channelName := "log"
	
		fmt.Println("Listening for messages...")
	
		// The goroutine stops here until it recieves a message
		msg := <-in
	
		PublishMessageToQueueByName(ch, msg, channelName)
	}
}