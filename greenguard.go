package main

import (
	"log"

	"github.com/vacovsky/gogrow/ggapi"
	"github.com/vacovsky/gogrow/ggdata"
	"github.com/vacovsky/gogrow/ggmodels"
	"github.com/vacovsky/gogrow/ggservice"
	"github.com/vacovsky/gogrow/ggservice/camera"
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
