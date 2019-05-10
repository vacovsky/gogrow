package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vacovsky/gogrow/ggdata"
)

func Dump(w http.ResponseWriter, r *http.Request) {
	dump := ggdata.GetLatestDataDump()
	blob, err := json.Marshal(&dump)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, string(blob))
}
