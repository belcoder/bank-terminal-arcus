package rabbit

import (
	"../../../pkg/logger"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

var AmqpConnection *amqp.Connection = nil
var AmqpChanel *amqp.Channel = nil
var Context context.Context
var Queue map[string]amqp.Queue

func Connect() bool {
	conn, err := amqp.Dial("amqp://belcoder:qwer1331@10.3.0.102:5672")
	if err != nil {
		logger.New("rabbit:Connect", err.Error())
		return false
	}

	AmqpConnection = conn
	return true
}

func OpenChanel() bool {
	if AmqpConnection == nil {
		return false
	}

	ch, err := AmqpConnection.Channel()
	if err != nil {
		logger.New("rabbit:OpenChanel", err.Error())
		return false
	}

	AmqpChanel = ch
	return true
}

func QueueDeclare(name string) bool {
	if AmqpChanel == nil {
		return false
	}

	queue, err := AmqpChanel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		logger.New("rabbit:QueueDeclare", err.Error())
		return false
	}

	Queue[name] = queue
	return true
}

func PublishWithContext(queueName string, body string) {
	if AmqpChanel == nil {
		return
	}

	err := AmqpChanel.PublishWithContext(Context,
		"",                    // exchange
		Queue[queueName].Name, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		logger.New("rabbit:PublishWithContext", err.Error())
	}
}
