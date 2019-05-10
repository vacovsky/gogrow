package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetFan() ggmodels.Fan {
	var f ggmodels.Fan
	Service().Last(&f)
	return f
}
