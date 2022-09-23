package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	firecracker "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

func getFirecrackerConfig(vmmId string) (firecracker.Config, error) {
	socketPath := getSocketPath(vmmId)
	return firecracker.Config{
		SocketPath:      socketPath,
		KernelImagePath: "../fire/build/hello-vmlinux.bin",
		LogPath:         fmt.Sprintf("%s.log", socketPath),
		Drives: []models.Drive{{
			DriveID:      firecracker.String("1"),
			PathOnHost:   firecracker.String("../fire/build/rootfs.ext4"),
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
