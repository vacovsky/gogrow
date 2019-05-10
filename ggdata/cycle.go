package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetGrowCycle() ggmodels.Cycle {
	var d ggmodels.Cycle
	Service().Last(&d)
	return d
}
