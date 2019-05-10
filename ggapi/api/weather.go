package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vacovsky/gogrow/ggservice/weather"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	result := weather.GetWeatherFromAPI()
	blob, err := json.Marshal(&result)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(blob))
}
