package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/vacovsky/gogrow/ggdata"
	"github.com/vacovsky/gogrow/ggmodels"
	"github.com/davecgh/go-spew/spew"
)

func Light(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var model ggmodels.Light
		err := json.NewDecoder(r.Body).Decode(&model)
		ggdata.Service().Save(&model)
		spew.Dump(model)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		result := ggmodels.Light{}
		blob, err := json.Marshal(&result)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		io.WriteString(w, string(blob))
	}
}
