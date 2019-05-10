package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vacovsky/gogrowggmodels"

	"github.com/vacovsky/gogrowggdata"
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
