package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetFilter() ggmodels.Filter {
	var d ggmodels.Filter
	Service().Last(&d)
	return d
}
