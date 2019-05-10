package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetGrowCycle() ggmodels.Cycle {
	var d ggmodels.Cycle
	Service().Last(&d)
	return d
}
