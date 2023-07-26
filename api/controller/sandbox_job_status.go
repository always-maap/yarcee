package controller

import (
	"api/database"
	"api/model"
	"fmt"
)

type JobStatus struct {
	ID           uint   `json:"id"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Error        string `json:"error"`
	StdErr       string `json:"stderr"`
	StdOut       string `json:"stdout"`
	ExecDuration int    `json:"exec_duration"`
	MemUsage     int64  `json:"mem_usage"`
}

func SandboxJobStatusHandler(jobStatus JobStatus) {
	fmt.Println("processing job status")
	sandbox := model.Sandbox{
		Id: jobStatus.ID,
	}

	database.DB.Model(&sandbox).Updates(model.Sandbox{
		Status:       jobStatus.Status,
		ExecDuration: jobStatus.ExecDuration,
		ExecMemUse:   int(jobStatus.MemUsage),
		StdOut:       jobStatus.StdOut,
		StdErr:       jobStatus.StdErr,
	})

}
