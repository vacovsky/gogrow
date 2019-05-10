package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetFertilization() ggmodels.FertilizationEvent {
	var f ggmodels.FertilizationEvent
	Service().Last(&f)
	return f
}
