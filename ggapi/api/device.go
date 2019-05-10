package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggmodels"

	"bitbucket.org/vacovsky/greenguard/ggdata"
)

func Device(w http.ResponseWriter, r *http.Request) {

	// timeScale, err := strconv.Atoi(r.Form.Get("timeScale"))
	// serial := string(r.Form.Get("serial"))

	// result := watrdevice.GetDeviceMetrics(account.ID, serial, timeScale)

	var dev ggmodels.DeviceState
	ggdata.Service().First(&dev)

	blob, err := json.Marshal(&dev)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w,
		"{ \"Device\":"+string(blob)+"}",
	)
}
