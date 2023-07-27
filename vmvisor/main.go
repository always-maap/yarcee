package main

import (
	"context"
	"encoding/json"
	"fmt"

	logrus "github.com/sirupsen/logrus"
)

type sandboxJob struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

var (
	job_q jobChan
)

func main() {
	defer deleteVMMSockets()
	vmmCtx, vmmCancel := context.WithCancel(context.Background())
	defer vmmCancel()

	warmVMs := make(chan runningFirecracker, 10)

	go fillVMPool(vmmCtx, warmVMs)

	logrus.SetReportCaller(true)

	job_q = newJobQueue("amqp://guest:guest@localhost:5672/")
	defer job_q.conn.Close()
	defer job_q.ch.Close()

	forever := make(chan bool)
	go func() {
		for d := range job_q.jobs {
			var job sandboxJob
			fmt.Println(string(d.Body))
			err := json.Unmarshal([]byte(d.Body), &job)
			if err != nil {
				logrus.WithError(err).Error("Received invalid job")
				continue
			}

			logrus.Info("Received a job ", job)
			job.run(vmmCtx, warmVMs)
		}
	}()

	logrus.Info(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
