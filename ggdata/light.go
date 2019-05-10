package ggdata

import "github.com/vacovsky/gogrowggmodels"

func GetLight() ggmodels.Light {
	var l ggmodels.Light
	Service().Last(&l)
	return l
}
