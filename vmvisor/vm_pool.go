package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

func fillVMPool(ctx context.Context, WarmVMs chan<- runningFirecracker) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			vm, err := createAndStartVM(ctx)
			if err != nil {
				continue
			}

			log.WithField("ip", vm.ip).Info("New VM created and started")

			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			err = waitForVMToBoot(ctx, vm.ip)
			if err != nil {
				log.WithError(err).Info("VM not ready yet")
				vm.vmmCancel()
				continue
			}

			// Add the new microVM to the pool.
			// If the pool is full, this line will block until a slot is available.
			WarmVMs <- *vm
		}
	}
}
