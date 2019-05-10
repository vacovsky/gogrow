package ggdata

import "github.com/vacovsky/gogrow/ggmodels"

func GetLight() ggmodels.Light {
	var l ggmodels.Light
	Service().Last(&l)
	return l
}
