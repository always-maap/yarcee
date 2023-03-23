package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type runningVM struct {
	
}

func main() {
	ctx := context.Background()
	vmmCtx, vmmCancel := context.WithCancel(ctx)
	defer vmmCancel()

	log.SetReportCaller(true)

	createAndStartVmm(vmmCtx)
}
