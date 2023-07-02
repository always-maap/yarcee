package main

import (
	"context"
	"encoding/json"

	logrus "github.com/sirupsen/logrus"
)

type sandboxJob struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

var (
	q jobQueue
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

	q = newJobQueue("amqp://guest:guest@localhost:5672/")
	defer q.conn.Close()
	defer q.ch.Close()

	forever := make(chan bool)
	go func() {
		for d := range q.jobs {
			var job sandboxJob
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
