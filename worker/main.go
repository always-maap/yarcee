package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	vmmCtx, vmmCancel := context.WithCancel(ctx)
	defer vmmCancel()

	log.SetReportCaller(true)

	createAndStartVmm(vmmCtx)
}
