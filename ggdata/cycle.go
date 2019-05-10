package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetGrowCycle() ggmodels.Cycle {
	var d ggmodels.Cycle
	Service().Last(&d)
	return d
}
