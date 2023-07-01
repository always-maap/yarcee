package main

import (
	"github.com/sirupsen/logrus"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}

type jobQueue struct {
	ch    *amqp.Channel
	conn  *amqp.Connection
	jobsQ amqp.Queue
	jobs  <-chan amqp.Delivery
}

func newJobQueue(amqpURL string) jobQueue {
	conn, err := amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

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

	sandboxQ, err := ch.QueueDeclare(
		"sandbox_queue", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		sandboxQ.Name,    // queue name
		"sandbox_job_rk", // routing key
		"sandbox_job_ex", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	jobs, err := ch.Consume(
		sandboxQ.Name, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	return jobQueue{
		ch,
		conn,
		sandboxQ,
		jobs,
	}

}
