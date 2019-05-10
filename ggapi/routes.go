package ggapi

import (
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggapi/api"
)

var unprotectedRoutes = map[string]func(http.ResponseWriter, *http.Request){
	// root

	"/dump":          api.Dump,
	"/environment":   api.Environment,
	"/fertilization": api.Fertilization,
	"/device":        api.Device,
	"/light":         api.Light,
	"/fan":           api.Fan,
	"/cycle":         api.Cycle,
	"/weather":       api.Weather,
	"/water":         api.Water,

	"/chart/temp": api.ChartTemperature,
	"/chart/hum":  api.ChartHumidity,

	// "/dashboard": dashboard,
	// "/digest":    digest,
}
