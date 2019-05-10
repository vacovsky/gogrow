package ggdata

import (
	"github.com/vacovsky/gogrow/ggmodels"
)

func GetLatestWeather() ggmodels.Weather {
	var w ggmodels.Weather
	Service().Last(&w)
	return w
}
