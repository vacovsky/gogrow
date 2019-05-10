package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func Fertilization(w http.ResponseWriter, r *http.Request) {

	// timeScale, err := strconv.Atoi(r.Form.Get("timeScale"))
	// serial := string(r.Form.Get("serial"))

	// result := watrdevice.GetDeviceMetrics(account.ID, serial, timeScale)

	result := ggmodels.Environment{}

	envblob, err := json.Marshal(&result)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w,
		"{ \"Environment\":"+string(envblob)+"}",
	)
}
