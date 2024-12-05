package airsend

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// types
type ConfigAirSend struct {
	Ip       string
	Password string
	BaseUrl  string
}

type Device struct {
	Id      int16
	Name    string
	Type    int16
	Pid     int16
	Attr    int16
	Opt     int16
	Counter int16
}

type AirSendClient struct {
	config     ConfigAirSend
	httpClient *http.Client
}

type AirSendClientInterface interface {
	GetDevices() ([]Device, error)
	ActionDevices()
}

func NewAirSendClient() (*AirSendClient, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &AirSendClient{
		config: ConfigAirSend{
			Ip:       os.Getenv("AIRSEND_IP"),
			Password: os.Getenv("AIRSEND_PASSWORD"),
			BaseUrl:  os.Getenv("AIRSEND_BASEURL"),
		},
		httpClient: &http.Client{},
	}, nil
}

func (c *AirSendClient) GetDevices() ([]Device, error) {
	log.Println("GetDevice")
	return []Device{}, nil
}

func (c *AirSendClient) ActionDevice(action string, deviceId int16) {
	log.Println("ActionDevice")
}
