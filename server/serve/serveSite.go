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

func Loadsites() error {
	sitecount, _ := ioutil.ReadDir(util.Sitesdir)
	sites = make(map[string][]byte, len(sitecount))

	for _, site := range sitecount {
		tmpsite, err := ioutil.ReadFile(util.Sitesdir + "/" + site.Name())
		if err != nil {
			util.Err(util.SERVE, err, true, "Error loading site")
			return err
		}
		sites[site.Name()] = tmpsite
		util.Log(util.SERVE, "Loaded site:", site.Name())
	}

	util.Log(util.MAIN, "All sites loaded")
	return nil
}

func getSite(name string) ([]byte, error) {
	if sites[name] != nil {
		return sites[name], nil
	} else {
		return nil, errors.New("no site loaded for: " + name)
	}
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

		msg, err := getSite("index.html")

		if err != nil {
			util.Err(util.SERVE, err, true, "Error getting site")
			w.WriteHeader(http.StatusInternalServerError)
			msg = []byte("Error serving site")
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Sent site successfully")
		}
	})

	// Other Sites
	rout.HandleFunc("/{site}", func(w http.ResponseWriter, r *http.Request) {
		msg, err := getSite(r.URL.Path[1:])

		if r.URL.Path[1:] == "api" {
			util.Err(util.SERVE, err, true, "API requested")
			w.WriteHeader(http.StatusForbidden)
			msg = []byte("Site Forbidden")
		} else if err != nil {
			util.Err(util.SERVE, err, true, "Error getting site")
			w.WriteHeader(http.StatusNotFound)
			msg = []byte("Site not found")
		} else {
			switch strings.Split(r.URL.Path[1:], ".")[1] {
			case "css":
				w.Header().Set("Content-Type", "text/css; charset=utf-8")
			case "ico":
				w.Header().Set("Content-Type", "image/x-icon")
			}
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		}
	})
}
