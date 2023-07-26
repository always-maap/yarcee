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
	//deleteVMMSockets()
	ctx := context.Background()
	vmmCtx, vmmCancel := context.WithCancel(ctx)
	defer vmmCancel()

	logrus.SetReportCaller(true)

	vm, err := createAndStartVM(vmmCtx)
	if err != nil {
		panic("failed to run vm")
	}
	logrus.Info(vm.ip)

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
			job.run(ctx, *vm)
		}
	}()

	logrus.Info(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
