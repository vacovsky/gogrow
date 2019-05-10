package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetDeviceState() ggmodels.DeviceState {
	var d ggmodels.DeviceState
	Service().Last(&d)
	return d
}
