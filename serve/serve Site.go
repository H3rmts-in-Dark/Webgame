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
	site, err = ioutil.ReadFile("test.html")
	if err != nil {
		util.Log("Serve", "err:", err)
		return err
	}
	util.Log("Serve", "loaded Site")
	return nil
}

func getSite() ([]byte, error) {
	if site != nil {
		return site, nil
	} else {
		return nil, errors.New("site not loaded")
	}
}

/*
CreateServe

Registers '/' handle to server the html site
*/
func CreateServe(rout *mux.Router) {
	rout.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		
		util.Log("Serve", "received request ", r.RemoteAddr)
		
		msg, err := getSite()
		
		if err != nil {
			util.Log("Serve", "err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			msg = []byte("Error serving site")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		}
		
		_, err = w.Write(msg)
		if err != nil {
			util.Log("API", "err in sending:", err)
		} else {
			util.Log("API", "send site")
		}
	})
}
