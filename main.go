package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metal "github.com/equinix-labs/metal-go/metal/v1"
)

var (
	version = "dev"
	warn    = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	client  *metal.APIClient
)

const (
	uaFmt = "gh-action-metal-deleter/%s %s"
)

func main() {
	authToken := os.Getenv("INPUT_AUTHTOKEN")
	projectID := os.Getenv("INPUT_PROJECTID")
	runnerName := os.Getenv("INPUT_RUNNERNAME")

	if authToken == "" {
		log.Fatal("You must provide an auth token in `with.userToken`")
	}

	if projectID == "" {
		log.Fatal("You must specify a project ID in `with.projectID`")
	}

	if runnerName == "" {
		log.Fatal("You must specify a runner name in `with.runnerName`")
	}

	config := metal.NewConfiguration()
	config.AddDefaultHeader("X-Auth-Token", authToken)
	config.UserAgent = fmt.Sprintf(uaFmt, version, config.UserAgent)
	client = metal.NewAPIClient(config)

	// What to delete?
	// - Devices
	// - ProjectVirtualNetworks
	// - Project
	// - TODO(displague) SpotMarkets
	// - TODO(displague) VolumeAttachments
	// - TODO(displague) Volumes (are these project specific?)

	devices, err := getAllProjectDevices(projectID)
	if err != nil {
		warn.Println("Could not list devices", err)
	}

	for _, device := range devices {
		hostName := device.GetHostname()
		if hostName == runnerName {
			fmt.Println("Deleting device", hostName)
			if _, err = client.DevicesApi.DeleteDevice(context.Background(), device.GetId()).Execute(); err != nil {
				warn.Println("Could not delete device", err)
			}
		}
	}
}

func getAllProjectDevices(projectID string) ([]metal.Device, error) {
	var devices []metal.Device
	var page int32 = 1

	for {
		devicePage, _, err := client.DevicesApi.FindProjectDevices(context.Background(), projectID).Page(page).Execute()
		if err != nil {
			return nil, err
		}

		devices = append(devices, devicePage.Devices...)
		if devicePage.Meta.GetLastPage() > devicePage.Meta.GetCurrentPage() {
			page++
			continue
		}

		return devices, nil
	}
}
