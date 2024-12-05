package main

import (
	service_airsend "MonitoringVoletsWithTime/src/domaine/service"
	"MonitoringVoletsWithTime/src/utils"
	"fmt"
	"github.com/nathan-osman/go-sunrise"
	"log"
	"time"
)

func main() {
	switterSerice := service_airsend.AirSendService{}

	// Boucle principale qui tourne toutes les minutes
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Println("Start loop")

		select {
		case <-ticker.C:
			now := time.Now()

			fmt.Println("Now: ", now)

			sunrise, sunset := sunrise.SunriseSunset(
				48.8566, 3522,
				now.Year(), now.Month(), now.Day())

			if utils.IsTimeToAct(now, sunrise, "sunrise") {
				log.Println("Time for sunrise!!")
				switterSerice.OpenShutter()
			}

			if utils.IsTimeToAct(now, sunset, "sunset") {
				log.Println("It's sunset time!")
				switterSerice.CloseShutter()
			}
		}

		fmt.Println("Stop loop")
	}
}
