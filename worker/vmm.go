package main

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"time"

	firecracker "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/google/uuid"
	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

type runningFirecracker struct {
	vmmCancel context.CancelFunc
	ip        net.IP
}

func createAndStartVmm(ctx context.Context) (*runningFirecracker, error) {
	// create vmmId with uuid
	vmmId := uuid.New().String()

	firecrackerCfg, err := getFirecrackerConfig(vmmId)

	if err != nil {
		return nil, err
	}

	logger := log.New()

	machineOpts := []firecracker.Opt{
		firecracker.WithLogger(log.NewEntry(logger)),
	}

	firecrackerBin, err := exec.LookPath("firecracker")
	if err != nil {
		fmt.Println("firecracker not found in PATH")
		return nil, err
	}

	cmd := firecracker.VMCommandBuilder{}.
		WithBin(firecrackerBin).
		WithSocketPath(firecrackerCfg.SocketPath).
		Build(ctx)
	machineOpts = append(machineOpts, firecracker.WithProcessRunner(cmd))

	vmmCtx, vmmCancel := context.WithCancel(ctx)

	machine, err := firecracker.NewMachine(vmmCtx, firecrackerCfg, machineOpts...)

	if err != nil {
		vmmCancel()
		return nil, fmt.Errorf("failed creating machine: %s", err)
	}

	if err := machine.Start(vmmCtx); err != nil {
		vmmCancel()
		return nil, fmt.Errorf("failed starting machine: %s", err)
	}

	return &runningFirecracker{
		vmmCancel: vmmCancel,
		ip:        machine.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr.IP,
	}, nil
}

func waitForVMToBoot(ctx context.Context, ip net.IP) error {
	// Query the agent until it provides a valid response
	req.SetTimeout(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			// Timeout
			return ctx.Err()
		default:
			res, err := req.Get("http://" + ip.String() + ":8080/health")
			if err != nil {
				log.WithError(err).Info("VM not ready yet")
				time.Sleep(time.Second)
				continue
			}

			if res.Response().StatusCode != 200 {
				time.Sleep(time.Second)
				log.Info("VM not ready yet")
			} else {
				log.WithField("ip", ip).Info("VM agent ready")
				return nil
			}
			time.Sleep(time.Second)
		}

	}
}
