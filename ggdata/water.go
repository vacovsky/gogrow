package ggdata

import (
	"log"

	"github.com/vacovsky/gogrow/ggmodels"
)

func GetWaterEvent() ggmodels.WateringEvent {
	var d ggmodels.WateringEvent
	Service().Last(&d)
	return d
}

func SaveWaterEvent(we ggmodels.WateringEvent) {
	log.Println("watering event saved")
	Service().Save(&we)
}
