package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggservice/weather"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	result := weather.GetWeatherFromAPI()
	blob, err := json.Marshal(&result)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(blob))
}
