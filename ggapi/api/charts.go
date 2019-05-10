package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/vacovsky/gogrow/ggdata"
)

func ChartTemperature(w http.ResponseWriter, r *http.Request) {

	result := ggdata.LoadTempChart()

	blob, err := json.Marshal(&result)

	if err != nil {
		log.Println("ERROR", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	io.WriteString(w, string(blob))
}

func ChartHumidity(w http.ResponseWriter, r *http.Request) {

	result := ggdata.LoadHumidityChart()

	blob, err := json.Marshal(&result)

	if err != nil {
		log.Println("ERROR", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	io.WriteString(w, string(blob))
}
