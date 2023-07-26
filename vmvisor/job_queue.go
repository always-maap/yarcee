package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}

type jobChan struct {
	ch    *amqp.Channel
	conn  *amqp.Connection
	jobsQ amqp.Queue
	jobs  <-chan amqp.Delivery
}

type jobStatus struct {
	ID           uint   `json:"id"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Error        string `json:"error"`
	StdErr       string `json:"stderr"`
	StdOut       string `json:"stdout"`
	ExecDuration int    `json:"exec_duration"`
	MemUsage     int64  `json:"mem_usage"`
}

func newJobQueue(amqpURL string) jobChan {
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
	failOnError(err, "Failed to declare sandbox job exchange")

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

	return jobChan{
		ch,
		conn,
		sandboxQ,
		jobs,
	}
}

func (q jobChan) setjobStatus(ctx context.Context, job sandboxJob, status string, res sandboxExecRes) error {
	log.WithField("status", status).Info("Set job status")
	jobStatus := &jobStatus{
		ID:      job.ID,
		Status:  status,
		Message: res.Message,
		StdErr:  "",
		StdOut:  "",
		//Error:   res.Error,
	}
	fmt.Printf("%+v", jobStatus)
	b, err := json.Marshal(jobStatus)
	if err != nil {
		return err
	}
	err = q.ch.Publish(
		"sandbox_job_ex",         // exchange
		"sandbox_jobs_status_rk", // routing key
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	return err
}

func (q jobChan) setJobReceived(ctx context.Context, job sandboxJob) error {
	return q.setjobStatus(ctx, job, "received", sandboxExecRes{})
}

func (q jobChan) setJobRunning(ctx context.Context, job sandboxJob) error {
	return q.setjobStatus(ctx, job, "running", sandboxExecRes{})
}

func (q jobChan) setJobFailed(ctx context.Context, job sandboxJob, res sandboxExecRes) error {
	return q.setjobStatus(ctx, job, "failed", res)
}

func (q jobChan) setJobResult(ctx context.Context, job sandboxJob, res sandboxExecRes) error {
	jobStatus := &jobStatus{
		ID:           job.ID,
		Status:       "done",
		Message:      res.Message,
		StdErr:       res.StdErr,
		StdOut:       res.StdOut,
		ExecDuration: res.ExecDuration,
		MemUsage:     res.MemUsage,
		//Error:        res.Error,
	}
	fmt.Printf("%+v", jobStatus)
	log.WithField("jobStatus", jobStatus).Info("Set job result")

	b, err := json.Marshal(jobStatus)
	if err != nil {
		return err
	}
	err = q.ch.Publish(
		"sandbox_job_ex",        // exchange
		"sandbox_job_status_rk", // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	return err
}
