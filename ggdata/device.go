package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetDeviceState() ggmodels.DeviceState {
	var d ggmodels.DeviceState
	Service().Last(&d)
	return d
}
