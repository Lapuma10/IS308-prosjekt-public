package rbMQ

import (
	"log"

	"github.com/streadway/amqp"
)

// might be moved to seperate file
const (
	RBMQ_URL = "amqp://guest:guest@rabbit1:5672/"
)

func PublishMessageToQueueByName(ch *amqp.Channel, message string, name string) {
	q := GetQueueInChannelByName(ch, name)

	body := message

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

// Use this to subscribe to queues
func GetChannelDataByName(ch *amqp.Channel, name string) <-chan amqp.Delivery {
	q := GetQueueInChannelByName(ch, name)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Channel that will recieve messages from rabbitMQ
	return msgs
}

func GetQueueInChannelByName(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}
func GetChannelFromConnection(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func GetConnection() *amqp.Connection {
	conn, err := amqp.Dial(RBMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
