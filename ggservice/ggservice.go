package ggservice

import (
	"log"
	"os"
	"strconv"
	"time"

	"bitbucket.org/vacovsky/greenguard/ggdata"
	"bitbucket.org/vacovsky/greenguard/ggservice/systemhealth"
	"bitbucket.org/vacovsky/greenguard/ggservice/weather"
)

var pollingInterval int

func init() {
	var err error
	pollingInterval, err = strconv.Atoi(os.Getenv("GG_POLLING_INTERVAL"))
	if err != nil {
		log.Println(err)
	}
}

// Start the hardware-interacting service
func Start() {

	go func() {
		for {
			saveDeviceInfo()
			time.Sleep(time.Second * time.Duration(pollingInterval))
		}
	}()

	go func() {
		for {
			saveWeatherInfo()
			time.Sleep(time.Minute * time.Duration(30))
		}
	}()

}

func saveWeatherInfo() {
	weather := weather.GetWeatherFromAPI()
	ds := ggdata.Service()
	ds.Save(&weather)

}

func saveDeviceInfo() {
	data := systemhealth.GatherDeviceInfo()
	ds := ggdata.Service()
	ds.Save(&data)
}
