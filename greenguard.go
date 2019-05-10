package main

import (
	"log"

	"github.com/vacovsky/gogrowggapi"
	"github.com/vacovsky/gogrowggdata"
	"github.com/vacovsky/gogrowggmodels"
	"github.com/vacovsky/gogrowggservice"
	"github.com/vacovsky/gogrowggservice/camera"
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
