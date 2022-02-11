package serve

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"Server/util"
)

var root dir

func LoadSites() error {
	util.Log(util.SERVE, "Loading Sites into Cache")
	start := time.Now()
	var size uint64
	var count uint32
	root, size, count = loadDir(util.GetConfig().SitesDir)
	util.Log(util.SERVE, fmt.Sprintf("All files (%d) loaded in %s; Size %dMB", count, time.Since(start), size/1048576))

	return nil
}

type dir struct {
	files map[string][]byte
	dirs  map[string]dir
}

func loadDir(name string) (dir, uint64, uint32) {
	siteCount, err := ioutil.ReadDir(name)
	if err != nil {
		util.Err(util.SERVE, err, false, "Error reading directory", name)
		return dir{}, 0, 0
	}
	dir := dir{map[string][]byte{}, map[string]dir{}}

	var size uint64 = 0
	var count uint32 = 0
	for _, site := range siteCount {
		if info, _ := os.Stat(name + "/" + site.Name()); info.IsDir() {
			dr, s, c := loadDir(name + "/" + site.Name())
			size += s
			count += c
			dir.dirs[site.Name()] = dr
			util.Debug(util.SERVE, "Loaded directory in cache", fmt.Sprintf("%s/%s", name, site.Name()))
		} else {
			tmpSite, err := ioutil.ReadFile(name + "/" + site.Name())
			if err != nil {
				util.Err(util.SERVE, err, false, "Error loading site", fmt.Sprintf("%s/%s", name, site.Name()))
			} else {
				size += uint64(len(tmpSite))
				count++
				dir.files[site.Name()] = tmpSite
				util.Debug(util.SERVE, "Loaded site in cache", fmt.Sprintf("%s/%s", name, site.Name()))
			}
		}
	}
	return dir, size, count
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

func getSite(path string, host string) (site []byte, code int, err error) {
	code, err = 202, nil

	if util.GetConfig().Cache {
		for _, forbidden := range util.GetConfig().Forbidden.Endpoints {
			if strings.HasPrefix(path, forbidden+"/") || path == forbidden {
				site, code = GetErrorSite(Forbidden, host, path)
				err = errors.New(path + " Forbidden by Endpoints " + forbidden)
				return
			}
		}

		for _, forbidden := range util.GetConfig().Forbidden.Regex {
			match, er := regexp.Match(forbidden, []byte(path))
			if er != nil {
				site, code = GetErrorSite(InternalServerError, host, path, fmt.Sprintf("Error checking forbidden regex"))
				err = er
				return
			}
			if match {
				site, code = GetErrorSite(Forbidden, host, path)
				err = errors.New(path + " Forbidden by regex " + forbidden)
				return
			}
		}

		pathSplit := strings.Split(path, "/")[1:]

		depth := len(pathSplit)

		dir := root
		for i := 0; i < depth-1; i++ {
			dir = dir.dirs[pathSplit[i]]
		}
		site = dir.files[pathSplit[depth-1]]
		if site == nil {
			if _, ok := dir.dirs[pathSplit[depth-1]]; ok {
				site, code = GetErrorSite(NotFound, host, path, fmt.Sprintf("%s is no file, but a directory", pathSplit[depth-1]))
				err = errors.New(fmt.Sprintf("no site data for: %v", pathSplit))
				return
			}
			site, code = GetErrorSite(NotFound, host, path)
			err = errors.New(fmt.Sprintf("no site data for: %s", pathSplit))
			return
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
		msg, code, err := getSite(r.URL.Path, r.Host)

		if err != nil {
			util.Err(util.SERVE, err, false, fmt.Sprintf("Error getting site %s", r.URL.Path))
			w.WriteHeader(code)
		} else {
			fileSplit := strings.Split(r.URL.Path[1:], ".")
			filetype := fileSplit[len(fileSplit)-1]
			if val, exists := util.GetConfig().ContentTypes[filetype]; exists == true {
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
