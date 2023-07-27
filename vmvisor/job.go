package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type sandboxExecReq struct {
	ID       uint   `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type sandboxExecRes struct {
	Message      string `json:"message"`
	StdErr       string `json:"stderr"`
	StdOut       string `json:"stdout"`
	ExecDuration int    `json:"exec_duration"`
	MemUsage     int64  `json:"mem_usage"`
}

func (job sandboxJob) run(ctx context.Context, warmVMs <-chan runningFirecracker) {
	log.WithField("job", job).Info("Handling job")

	vm := <-warmVMs

	// Defer cleanup of VM and VMM
	go func() {
		defer vm.vmmCancel()
		vm.machine.Wait(vm.vmmCtx)
	}()
	defer vm.shutDown()

	var reqJSON []byte

	reqJSON, err := json.Marshal(sandboxExecReq{
		ID:       job.ID,
		Language: job.Language,
		Code:     job.Code,
	})
	if err != nil {
		log.WithError(err).Error("Failed to marshal JSON request")
		return
	}

	if err != nil {
		log.WithError(err).Error("Could not set job running")
		return
	}

	var httpRes *http.Response
	var sandboxRes sandboxExecRes

	httpRes, err = http.Post("http://"+vm.ip.String()+":8080/exec", "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		log.WithError(err).Error("Failed to request execution to agent")
		return
	}
	json.NewDecoder(httpRes.Body).Decode(&sandboxRes)
	log.WithField("result", sandboxRes).Info("Job execution finished")
	if httpRes.StatusCode != 200 {
		log.WithFields(log.Fields{
			"httpRes":  httpRes,
			"agentRes": sandboxRes,
			"reqJSON":  string(reqJSON),
		}).Error("Failed to compile and run code")
		return
	}

	err = job_q.setJobResult(ctx, job, sandboxRes)
	if err != nil {
		job_q.setJobFailed(ctx, job, sandboxRes)
	}
}
