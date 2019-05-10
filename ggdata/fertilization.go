package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetFertilization() ggmodels.FertilizationEvent {
	var f ggmodels.FertilizationEvent
	Service().Last(&f)
	return f
}
