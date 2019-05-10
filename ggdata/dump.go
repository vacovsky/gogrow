package ggdata

import (
	"encoding/json"

	"github.com/vacovsky/gogrow/ggmodels"
)

// GetLatestEnvironmentCheck : return to the called the lastest environment DB entry
func GetLatestDataDump() ggmodels.DataDump {

	env := GetLatestEnvironmentCheck()
	light := GetLight()
	fan := GetFan()
	fert := GetFertilization()
	dev := GetDeviceState()
	weather := GetLatestWeather()
	filter := GetFilter()
	cycle := GetGrowCycle()
	water := GetWaterEvent()

	// envblob, err := json.Marshal(&env)
	// lightblob, err := json.Marshal(&light)
	// fanblob, err := json.Marshal(&fan)
	// fertblob, err := json.Marshal(&fert)
	// devblob, err := json.Marshal(&dev)
	// weatherblob, err := json.Marshal(&weather)
	// filterblob, err := json.Marshal(&filter)

	dump := ggmodels.DataDump{
		Environment:   env,
		Light:         light,
		Fan:           fan,
		Device:        dev,
		Fertilization: fert,
		Weather:       weather,
		Filter:        filter,
		Cycle:         cycle,
		Water:         water,
	}
	json.Marshal(&fan)

	return dump
}
