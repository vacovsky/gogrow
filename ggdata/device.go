package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetDeviceState() ggmodels.DeviceState {
	var d ggmodels.DeviceState
	Service().Last(&d)
	return d
}
