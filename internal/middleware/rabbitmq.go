package middleware

import (
	"NGB-SE/internal/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var senderChannel *amqp.Channel

type Message struct {
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
	TargetUid   string `json:"targetUid"`
}

func RabbitMQInit() {
	util.MakeInfoLog("Initializing RabbitMQ middleware")
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
		"",
		false,
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

func sendByteToQueue(databyte []byte) error {
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
	return err
}

func SendNotification(content, targetUid string) error {
	var newMsg Message
	newMsg.Body = content
	newMsg.ContentType = "Notification"
	newMsg.TargetUid = targetUid
	fmt.Println(newMsg.TargetUid)
	databyte, err := json.Marshal(&newMsg)
	if err != nil {
		util.MakeInfoLog("Failed whe marshaling")
	}
	err = sendByteToQueue(databyte)
	return err
}
