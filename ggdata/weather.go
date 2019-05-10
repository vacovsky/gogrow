package ggdata

import (
	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func GetLatestWeather() ggmodels.Weather {
	var w ggmodels.Weather
	Service().Last(&w)
	return w
}
