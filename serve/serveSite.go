package serve

import (
	"errors"
	"io/ioutil"
	"net/http"

	"Webgame/util"

	"github.com/gorilla/mux"
)

var site []byte

func Loadsites() error {
	var err error
	site, err = ioutil.ReadFile("resources/test.html")
	if err != nil {
		util.Err(util.SERVE, err, true, "Error loading site")
		return err
	}
	util.Log(util.SERVE, "Loaded site: resources/test.html")

	util.Log(util.MAIN, "All sites loaded")
	return nil
}

func getSite() ([]byte, error) {
	if site != nil {
		return site, nil
	} else {
		return nil, errors.New("no site loaded")
	}
}

/*
CreateServe

Registers a handle for '/' to serve the html site
*/
func CreateServe(rout *mux.Router) {
	rout.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		util.Log(util.SERVE, "Received request from", r.RemoteAddr)

		msg, err := getSite()

		if err != nil {
			util.Err(util.SERVE, err, true, "Error getting site")
			w.WriteHeader(http.StatusInternalServerError)
			msg = []byte("Error serving site")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Sent site successfully")
		}
	})
}
