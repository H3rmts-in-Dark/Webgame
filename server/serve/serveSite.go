package serve

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"Server/util"
)
import "unsafe"

var root dir

func LoadSites() error {
	util.Log(util.SERVE, "Loading Sites into Cache")
	start := time.Now()
	root = loadDir(util.GetConfig().SitesDir)
	util.Log(util.SERVE, "All sites loaded in", time.Since(start), " Size", unsafe.Sizeof(root))

	return nil
}

type dir struct {
	files map[string][]byte
	dirs  map[string]dir
}

func loadDir(name string) (data2 dir) {
	siteCount, err := ioutil.ReadDir(name)
	if err != nil {
		util.Err(util.SERVE, err, false, "Error reading directory", name)
		return dir{}
	}
	dir := dir{map[string][]byte{}, map[string]dir{}}

	for _, site := range siteCount {
		if info, _ := os.Stat(name + "/" + site.Name()); info.IsDir() {
			dir.dirs[site.Name()] = loadDir(name + "/" + site.Name())
			util.Debug(util.SERVE, "Loaded directory in cache", fmt.Sprintf("%s/%s", name, site.Name()))
		} else {
			tmpSite, err := ioutil.ReadFile(name + "/" + site.Name())
			if err != nil {
				util.Err(util.SERVE, err, false, "Error loading site", fmt.Sprintf("%s/%s", name, site.Name()))
			} else {
				dir.files[site.Name()] = tmpSite
				util.Debug(util.SERVE, "Loaded site in cache", fmt.Sprintf("%s/%s", name, site.Name()))
			}
		}
	}
	return dir
}

/*
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
			err = errors.New("no site data for: " + name)
		}
	} else {
		var tmpSite []byte
		tmpSite, err = ioutil.ReadFile(util.GetConfig().SitesDir + "/" + name)
		if err != nil {
			util.Err(util.SERVE, err, false, "Error loading site")
			site, code = GetErrorSite(NotFound, host)
		} else {
			site = tmpSite
		}
	}
	return
}
*/

func getSite(path []string, host string) (site []byte, code int, err error) {
	code, err = 202, nil

	if util.GetConfig().Cache { // TODO implement forebidden
		depth := len(path)

		dir := root
		for i := 0; i < depth-1; i++ {
			dir = dir.dirs[path[i]]
		}
		site = dir.files[path[depth-1]]
		if site == nil {
			site, code = GetErrorSite(NotFound, host)
			err = errors.New(fmt.Sprintf("no site data for: %v", path))
		}

	} else {
		// TODO implement reading
	}

	return
}

/*
CreateServe

Registers a handle for '/' to serve the DefaultSite
*/
func CreateServe() http.HandlerFunc {

	fun := func(w http.ResponseWriter, r *http.Request) {
		msg, code, err := getSite(strings.Split(r.URL.Path, "/")[1:], r.Host)

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
		}
	}

	return fun
}

func stringInList(search string, list []string) bool {
	for _, val := range list {
		if val == search {
			return true
		}
	}
	return false
}
