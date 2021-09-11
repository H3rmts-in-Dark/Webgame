package serve

import (
	"errors"
	"io/ioutil"
	"net/http"

	"Webgame/util"

	"github.com/gorilla/mux"
)

var site []byte

func LoadSite() {
	var err error
	site, err = ioutil.ReadFile("resources/test.html")
	if err != nil {
		util.Log("SERVE", "Error: ", err)
		panic(err)
	}
	util.Log("SERVE", "Loaded site")
}

func getSite() ([]byte, error) {
	if site != nil {
		return site, nil
	} else {
		return nil, errors.New("No site loaded")
	}
}

/*
CreateServe:
Registers a handle for '/' to serve the html site
*/
func CreateServe(rout *mux.Router) {
	rout.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		util.Log("SERVE", "received request from ", r.RemoteAddr)

		msg, err := getSite()

		if err != nil {
			util.Log("SERVE", "Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			msg = []byte("Error serving site")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Log("API", "Error in sending:", err)
		} else {
			util.Log("API", "Sent site successfully")
		}
	})
}
