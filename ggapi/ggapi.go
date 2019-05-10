package ggapi

import (
	"fmt"
	"net/http"
	"os"

	"bitbucket.org/vacovsky/greenguard/ggservice/camera"
)

// Start starts the API web application
func Start() {

	for k, v := range unprotectedRoutes {
		http.HandleFunc(k, v)
	}

	fmt.Println("API Endpoints loaded...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		camera.TakeNewPicture()
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	panic(http.ListenAndServe(":"+os.Getenv("GG_API_PORT"), nil))
}

// func logTraffic() string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	return runtime.FuncForPC(pc).Name()
// }
