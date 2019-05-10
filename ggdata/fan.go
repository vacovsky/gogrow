package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetFan() ggmodels.Fan {
	var f ggmodels.Fan
	Service().Last(&f)
	return f
}
