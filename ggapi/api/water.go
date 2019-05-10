package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"bitbucket.org/vacovsky/greenguard/ggdata"

	"bitbucket.org/vacovsky/greenguard/ggmodels"
)

func Water(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var water ggmodels.WateringEvent
		err := json.NewDecoder(r.Body).Decode(&water)
		ggdata.Service().Save(&water)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		result := ggmodels.WateringEvent{}
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
