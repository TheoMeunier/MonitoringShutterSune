package service_airsend

import (
	"MonitoringVoletsWithTime/src/infrastructure/airsend"
	"log"
)

type AirSendService struct {
	airSendClient *airsend.AirSendClientInterface
}

func (c *AirSendService) OpenShutter() {
	client, _ := airsend.NewAirSendClient()
	devices, _ := client.GetDevices()

	for _, device := range devices {
		client.ActionDevice("ON", device.Id)
		log.Println("open shutter", device.Name)
	}
}

func (c *AirSendService) CloseShutter() {
	client, _ := airsend.NewAirSendClient()
	devices, _ := client.GetDevices()

	for _, device := range devices {
		client.ActionDevice("OFF", device.Id)
		log.Println("close shutter", device.Name)
	}
}
