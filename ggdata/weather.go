package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

func GetLatestWeather() ggmodels.Weather {
	var w ggmodels.Weather
	Service().Last(&w)
	return w
}
