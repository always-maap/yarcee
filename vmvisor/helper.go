package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	firecracker "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/firecracker-microvm/firecracker-go-sdk/client/models"
	logrus "github.com/sirupsen/logrus"
)

func getFirecrackerConfig(vmmId string) (firecracker.Config, error) {
	socketPath := getSocketPath(vmmId)
	return firecracker.Config{
		SocketPath:      socketPath,
		KernelImagePath: "../frontline/build/hello-vmlinux.bin",
		LogPath:         fmt.Sprintf("%s.log", socketPath),
		Drives: []models.Drive{{
			DriveID:      firecracker.String("1"),
			PathOnHost:   firecracker.String("/tmp/rootfs-" + vmmId + ".ext4"),
			IsRootDevice: firecracker.Bool(true),
			IsReadOnly:   firecracker.Bool(false),
			RateLimiter: firecracker.NewRateLimiter(
				// bytes/s
				models.TokenBucket{
					OneTimeBurst: firecracker.Int64(1024 * 1024), // 1 MiB/s
					RefillTime:   firecracker.Int64(500),         // 0.5s
					Size:         firecracker.Int64(1024 * 1024),
				},
				// ops/s
				models.TokenBucket{
					OneTimeBurst: firecracker.Int64(100),  // 100 iops
					RefillTime:   firecracker.Int64(1000), // 1s
					Size:         firecracker.Int64(100),
				}),
		}},
		NetworkInterfaces: []firecracker.NetworkInterface{{
			// Use CNI to get dynamic IP
			CNIConfiguration: &firecracker.CNIConfiguration{
				NetworkName: "fcnet",
				IfName:      "veth0",
			},
		}},
		MachineCfg: models.MachineConfiguration{
			VcpuCount:  firecracker.Int64(1),
			Smt:        firecracker.Bool(true),
			MemSizeMib: firecracker.Int64(256),
		},
	}, nil
}

func getSocketPath(vmmID string) string {
	filename := strings.Join([]string{
		".firecracker.sock",
		strconv.Itoa(os.Getpid()),
		vmmID,
	},
		"-",
	)
	dir := os.TempDir()

	return filepath.Join(dir, filename)
}

func deleteVMMSockets() {
	logrus.Info("Deleting VMM sockets")
	dir, err := ioutil.ReadDir(os.TempDir())
	if err != nil {
		logrus.WithError(err).Error("Failed to read directory")
	}
	for _, d := range dir {
		if strings.Contains(d.Name(), ".firecracker.sock-") {
			logrus.WithField("d", d.Name()).Debug("deleting")
			os.Remove(path.Join([]string{"tmp", d.Name()}...))
		}
	}
}
