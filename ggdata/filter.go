package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetFilter() ggmodels.Filter {
	var d ggmodels.Filter
	Service().Last(&d)
	return d
}
