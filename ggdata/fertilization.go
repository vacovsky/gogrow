package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetFertilization() ggmodels.FertilizationEvent {
	var f ggmodels.FertilizationEvent
	Service().Last(&f)
	return f
}
