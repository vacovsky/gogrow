package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetFan() ggmodels.Fan {
	var f ggmodels.Fan
	Service().Last(&f)
	return f
}
