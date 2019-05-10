package ggdata

import (
	"log"

	"bitbucket.org/vacovsky/greenguard/ggmodels"
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
