package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetFilter() ggmodels.Filter {
	var d ggmodels.Filter
	Service().Last(&d)
	return d
}
