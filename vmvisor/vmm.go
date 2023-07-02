package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"time"

	firecracker "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/google/uuid"
	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

type runningFirecracker struct {
	vmmCtx    context.Context
	vmmCancel context.CancelFunc
	vmmID     string
	machine   *firecracker.Machine
	ip        net.IP
}

func copy(src string, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	return err
}

func createAndStartVM(ctx context.Context) (*runningFirecracker, error) {
	// create vmmId with uuid
	vmmId := uuid.New().String()

	copy("../frontline/build/rootfs.ext4", "/tmp/rootfs-"+vmmId+".ext4")

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

	logger.Printf("Starting VMM with ID %s", &machine.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr)

	return &runningFirecracker{
		vmmCancel: vmmCancel,
		vmmCtx:    vmmCtx,
		vmmID:     vmmId,
		machine:   machine,
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

func (vm runningFirecracker) shutDown() {
	log.WithField("ip", vm.ip).Info("stopping")
	vm.machine.StopVMM()
	err := os.Remove(vm.machine.Cfg.SocketPath)
	if err != nil {
		log.WithError(err).Error("Failed to delete firecracker socket")
	}
	err = os.Remove("/tmp/rootfs-" + vm.vmmID + ".ext4")
	if err != nil {
		log.WithError(err).Error("Failed to delete firecracker rootfs")
	}
}
