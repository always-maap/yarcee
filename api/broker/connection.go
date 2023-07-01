package broker

import (
	"log"
	"os"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex
)

func GetChannel() *amqp.Channel {
	return ch
}

func Connect() error {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	mu.Lock()
	defer mu.Unlock()

	if conn != nil && ch != nil {
		return nil
	}

	var err error

	conn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		return err
	}

	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		"sandbox_job_ex", // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	return nil
}
