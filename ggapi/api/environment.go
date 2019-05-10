package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func Environment(w http.ResponseWriter, r *http.Request) {

	// timeScale, err := strconv.Atoi(r.Form.Get("timeScale"))
	// serial := string(r.Form.Get("serial"))

	// result := watrdevice.GetDeviceMetrics(account.ID, serial, timeScale)

	result := ggmodels.Environment{}

	blob, err := json.Marshal(&result)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, string(blob))
}

// http://api.openweathermap.org/data/2.5/forecast?id=524901&APPID={APIKEY}
