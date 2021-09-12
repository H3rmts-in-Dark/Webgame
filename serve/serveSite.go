package serve

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"Webgame/util"
	
	"github.com/gorilla/mux"
)

var site []byte

func Loadsite() error {
	var err error
	site, err = ioutil.ReadFile("resources/test.html")
	if err != nil {
		util.Err(util.SERVE, err,true, "Error loading Site")
		return err
	}
	util.Log(util.SERVE, "Loaded Site: resources/test.html")
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
		fmt.Println()
		util.Log(util.SERVE, "received request from ", r.RemoteAddr)
		
		msg, err := getSite()
		
		if err != nil {
			util.Err(util.SERVE,err,true, "Error getting Site")
			w.WriteHeader(http.StatusInternalServerError)
			msg = []byte("Error serving site")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		}
		
		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err,true, "Error Writing Response:")
		} else {
			util.Log(util.SERVE, "Sent site successfully")
		}
	})
}
