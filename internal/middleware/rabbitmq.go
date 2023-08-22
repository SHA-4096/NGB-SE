package middleware

import (
	"NGB-SE/internal/util"
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var senderChannel *amqp.Channel

func init() {
	//establish connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		util.MakeInfoLog("Failed to connect to RabbitMQ")
	}
	//create channel
	senderChannel, err = conn.Channel()
	if err != nil {
		util.MakeInfoLog("[RabbitMQ]Failed to open a channel")
	}
	//decleare queue
	queue, err := senderChannel.QueueDeclare(
		"userActivity",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		util.MakeInfoLog("[RabbitMQ]Failed to decleare a queue")

	} else {
		util.MakeInfoLog(fmt.Sprintf("[RabbitMQ]Queue %s Decleared", queue.Name))
	}
}

func SendByteToQueue(databyte []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := senderChannel.PublishWithContext(ctx,
		"",             // exchange
		"userActivity", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        databyte,
		})
	if err != nil {
		util.MakeInfoLog("[RabbitMQ]Failed to publish a message")
	} else {
		util.MakeInfoLog("[RabbitMQ]Message sent")
	}
}
