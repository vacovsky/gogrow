package ggdata

import "bitbucket.org/vacovsky/greenguard/ggmodels"

func GetLight() ggmodels.Light {
	var l ggmodels.Light
	Service().Last(&l)
	return l
}
