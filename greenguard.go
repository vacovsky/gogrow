package main

import (
	"log"

	"bitbucket.org/vacovsky/greenguard/ggapi"
	"bitbucket.org/vacovsky/greenguard/ggdata"
	"bitbucket.org/vacovsky/greenguard/ggmodels"
	"bitbucket.org/vacovsky/greenguard/ggservice"
	"bitbucket.org/vacovsky/greenguard/ggservice/camera"
)

var (
	bus = map[string]chan string{}
)

func init() {
	log.Println("Generating channel map")

	bus["service"] = make(chan string)
	bus["api"] = make(chan string)
	bus["data"] = make(chan string)
}

func main() {

	var ds ggmodels.DeviceState
	if ds != ds {
	}
	ggdata.MigrateDataSchema()
	go ggservice.Start()
	go camera.Start()
	ggapi.Start()
}
