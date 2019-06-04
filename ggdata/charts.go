package ggdata

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vacovsky/gogrow/ggmodels"
)

var (
	chartHours int
)

func init() {
	var err error
	chartHours, err = strconv.Atoi(os.Getenv("GG_CHART_HOURS"))
	if err != nil {
		log.Println(err)
	}
}

func LoadTempChart() ggmodels.Chart {

	data := []ggmodels.DeviceState{}
	weatherData := []ggmodels.Weather{}

	labels := []int64{}

	ambientTempStr := []string{}
	ambientTentTempStr := []string{}
	cpuTempStr := []string{}
	outsideTempStr := []string{}

	Service().Where("time_stamp > ?", time.Now().Add(-chartHours*time.Hour)).Find(&data).Order("time_stamp desc")
	Service().Where("time_stamp > ?", time.Now().Add(-chartHours*time.Hour)).Find(&weatherData).Order("time_stamp desc")

	var curTT float64
	var curAT float64
	for _, v := range data {
		labels = append(labels, v.TimeStamp.Unix()*1000)
		if v.TentTemperature != 0 {
			curTT = v.TentTemperature
		}
		if v.AmbientTemperature != 0 {
			curAT = v.AmbientTemperature
		}
		ambientTempStr = append(ambientTempStr, fmt.Sprintf("%.2f", tempConv(curAT)))
		ambientTentTempStr = append(ambientTentTempStr, fmt.Sprintf("%.2f", tempConv(curTT)))
		cpuTempStr = append(cpuTempStr, fmt.Sprintf("%.2f", tempConv(v.CPUTemperatureC)))
	}

	var curOT float64
	for _, v := range data {
		for _, o := range weatherData {
			if v.TimeStamp.Unix() > o.TimeStamp.Unix() {
				curOT = o.Temperature
			}
		}
		outsideTempStr = append(outsideTempStr, fmt.Sprintf("%.2f", curOT))
	}

	chart := ggmodels.Chart{
		Series: []string{"Tent", "Ambient", "Outside", "Device CPU"},
		Labels: labels,
		Data:   [][]string{ambientTentTempStr, ambientTempStr, outsideTempStr, cpuTempStr},
	}

	return chart
}

func LoadHumidityChart() ggmodels.Chart {

	// GG_CHART_HOURS
	data := []ggmodels.DeviceState{}
	weatherData := []ggmodels.Weather{}

	labels := []int64{}

	ambientHumStr := []string{}
	ambientTentHumStr := []string{}
	outsideCloudsStr := []string{}
	outsideHumStr := []string{}

	Service().Where("time_stamp > ?", time.Now().Add(-chartHours*time.Hour)).Find(&data).Order("time_stamp desc")
	Service().Where("time_stamp > ?", time.Now().Add(-chartHours*time.Hour)).Find(&weatherData).Order("time_stamp desc")

	var curTH float64
	var curAH float64
	for _, v := range data {
		labels = append(labels, v.TimeStamp.Unix()*1000)

		if v.AmbientHumidity != 0 {
			curAH = v.AmbientHumidity
		}
		if v.TentHumidity != 0 {
			curTH = v.TentHumidity
		}
		ambientHumStr = append(ambientHumStr, fmt.Sprintf("%.2f", curAH))
		ambientTentHumStr = append(ambientTentHumStr, fmt.Sprintf("%.2f", curTH))
	}

	var curOH int
	var curC int
	for _, v := range data {
		for _, o := range weatherData {
			if v.TimeStamp.Unix() > o.TimeStamp.Unix() {
				curOH = o.Humidity
				curC = o.Clouds
			}
		}
		outsideCloudsStr = append(outsideCloudsStr, strconv.Itoa(curC))
		outsideHumStr = append(outsideHumStr, strconv.Itoa(curOH))
	}

	chart := ggmodels.Chart{
		Series: []string{"Tent", "Ambient", "Outside", "Clouds"},
		Labels: labels,
		Data:   [][]string{ambientTentHumStr, ambientHumStr, outsideHumStr, outsideCloudsStr},
	}
	return chart
}

func tempConv(n float64) float64 {
	return (n * 1.8) + 32
	// Formula
	// (3°C × 1.8) + 32 = 37.4°F
}
