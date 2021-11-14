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

func LoadSites() error {
	siteCount, err := ioutil.ReadDir(util.Sitesdir)
	if err != nil {
		util.Err(util.SERVE, err, true, "Error reading sites directory")
		return err
	}

	sites = make(map[string][]byte, len(siteCount))

	for _, site := range siteCount {
		tmpSite, err := ioutil.ReadFile(util.Sitesdir + "/" + site.Name())
		if err != nil {
			util.Err(util.SERVE, err, true, "Error loading site")
			return err
		}
		sites[site.Name()] = tmpSite
		util.Log(util.SERVE, "Loaded site in cache:", site.Name())
	}

	util.Log(util.SERVE, "All sites loaded")
	return nil
}

func getSite(name string, host string) (site []byte, code int, err error) {
	code, err = 202, nil
	if stringInList(name, util.GetConfig().Forbidden) {
		site, code = GetErrorSite(Forbidden, host)
		err = errors.New("site " + name + " Forbidden")
		return
	}
	if util.GetConfig().Cache {
		if sites[name] != nil {
			site = sites[name]
		} else {
			site, code = GetErrorSite(NotFound, host)
			err = errors.New("no site loaded for: " + name)
		}
	} else {
		var tmpSite []byte
		tmpSite, err = ioutil.ReadFile(util.Sitesdir + "/" + name)
		if err != nil {
			util.Err(util.SERVE, err, false, "Error loading site")
			site, code = GetErrorSite(NotFound, host)
		} else {
			util.Log(util.SERVE, "Loaded site:", name)
			site = tmpSite
		}
	}
	return
}

/*
CreateServe

Registers a handle for '/' to serve the DefaultSite
*/
func CreateServe(rout *mux.Router) {
	// Main site
	rout.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()

		util.Log(util.SERVE, "Received request from", r.RemoteAddr, "for main site")

		msg, code, err := getSite(util.GetConfig().DefaultSite, r.Host)

		if err != nil {
			util.Err(util.SERVE, err, false, "Error getting main site")
			w.WriteHeader(code)
		} else {
			fileSplit := strings.Split(util.GetConfig().DefaultSite, ".")
			filetype := fileSplit[len(fileSplit)-1]
			if val, exists := util.GetConfig().Headers[filetype]; exists == true {
				w.Header().Set("Content-Type", val)
			}
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Send main site")
		}
	})

	// Other Sites
	rout.HandleFunc("/{site}", func(w http.ResponseWriter, r *http.Request) {
		msg, code, err := getSite(r.URL.Path[1:], r.Host)

		if err != nil {
			util.Err(util.SERVE, err, false, "Error getting site")
			w.WriteHeader(code)
		} else {
			fileSplit := strings.Split(r.URL.Path[1:], ".")
			filetype := fileSplit[len(fileSplit)-1]
			if val, exists := util.GetConfig().Headers[filetype]; exists == true {
				w.Header().Set("Content-Type", val)
			}
		}

		_, err = w.Write(msg)
		if err != nil {
			util.Err(util.SERVE, err, true, "Error writing response:")
		} else {
			util.Log(util.SERVE, "Sent site")
		}
	}).Methods("Get")
}

func stringInList(search string, list []string) bool {
	for _, val := range list {
		if val == search {
			return true
		}
	}
	return false
}
