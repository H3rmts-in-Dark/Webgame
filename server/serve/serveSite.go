package serve

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"Webgame/server/util"

	"github.com/gorilla/mux"
)

var sites map[string][]byte

var forebidden = map[string]bool{"api": true}

var headers = map[string]string{"css": "text/css; charset=utf-8", "ico": "image/x-icon"}

func LoadSites() error {
	sitecount, err := ioutil.ReadDir(util.Sitesdir)
	if err != nil {
		util.Err(util.SERVE, err, true, "Error reading sites directory")
		return err
	}

	sites = make(map[string][]byte, len(sitecount))

	for _, site := range sitecount {
		tmpsite, err := ioutil.ReadFile(util.Sitesdir + "/" + site.Name())
		if err != nil {
			util.Err(util.SERVE, err, true, "Error loading site")
			return err
		}
		sites[site.Name()] = tmpsite
		util.Log(util.SERVE, "Loaded site in cache:", site.Name())
	}

	util.Log(util.MAIN, "All sites loaded")
	return nil
}

func getSite(name string) (site []byte, code int, err error) {
	code, err = 202, nil
	if val, exits := forebidden[name]; exits == true && val == true {
		site, code = GetErrorSite(Forbidden)
		err = errors.New("site: " + name + " Forbidden")
		return
	}
	if util.GetConfig().Cache {
		if sites[name] != nil {
			site = sites[name]
		} else {
			site, code = GetErrorSite(NotFound)
			err = errors.New("no site loaded for: " + name)
		}
	} else {
		var tmpSite []byte
		tmpSite, err = ioutil.ReadFile(util.Sitesdir + "/" + name)
		if err != nil {
			util.Err(util.SERVE, err, false, "Error loading site")
			site, code = GetErrorSite(NotFound)
		} else {
			util.Log(util.SERVE, "Loaded site:", name)
			site = tmpSite
		}
	}
	return
}

/*
CreateServe

Registers a handle for '/' to serve the html site
*/
func CreateServe(rout *mux.Router) {
	// Main site
	rout.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		util.Log(util.SERVE, "Received request from", r.RemoteAddr, "for main site")

		msg, code, err := getSite("index.html")

		if err != nil {
			util.Err(util.SERVE, err, true, "Error getting site")
			w.WriteHeader(code)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Main Sent site successfully")
		}
	})

	// Other Sites
	rout.HandleFunc("/{site}", func(w http.ResponseWriter, r *http.Request) {
		msg, code, err := getSite(r.URL.Path[1:])

		if err != nil {
			util.Err(util.SERVE, err, false, "Error getting site")
			w.WriteHeader(code)
		} else {
			filetype := strings.Split(r.URL.Path[1:], ".")[1]
			if val, exists := headers[filetype]; exists == true {
				w.Header().Set("Content-Type", val)
			}
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Sent site successfully")
		}
	}).Methods("Get")
}
