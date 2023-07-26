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

func GetSandboxStatusJobsChan() <-chan amqp.Delivery {
	jobs, err := ch.Consume(
		"sandbox_status_queue", // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	failOnError(err, "Failed to register a consumer")

	return jobs
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
	failOnError(err, "Failed to declare sandbox job exchange")

	sandboxQ, err := ch.QueueDeclare(
		"sandbox_status_queue", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		sandboxQ.Name,           // queue name
		"sandbox_job_status_rk", // routing key
		"sandbox_job_ex",        // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	return nil
}
