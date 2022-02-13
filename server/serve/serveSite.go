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

	"Server/logging"
	"Server/util"
)

var root dir

func LoadSites() error {
	logging.Log(logging.SERVE, "Loading Sites into Cache")
	start := time.Now()
	var size uint64
	var count uint32
	root = loadDir(util.GetConfig().SitesDir, &size, &count)
	logging.Log(logging.SERVE, fmt.Sprintf("All files (%d) loaded in %s  Size:%dMB", count, time.Since(start), size/1048576))

	return nil
}

type dir struct {
	files map[string][]byte
	dirs  map[string]dir
}

func loadDir(name string, size *uint64, count *uint32) dir {
	siteCount, err := ioutil.ReadDir(name)
	if err != nil {
		logging.Err(logging.SERVE, err, false, "Error reading directory", name)
		return dir{}
	}
	dir := dir{map[string][]byte{}, map[string]dir{}}

	for _, site := range siteCount {
		if info, _ := os.Stat(name + "/" + site.Name()); info.IsDir() {
			dr := loadDir(name+"/"+site.Name(), size, count)
			dir.dirs[site.Name()] = dr
			logging.Debug(logging.SERVE, "Loaded directory in cache", fmt.Sprintf("%s/%s", name, site.Name()))
		} else {
			tmpSite, err := ioutil.ReadFile(name + "/" + site.Name())
			if err != nil {
				logging.Err(logging.SERVE, err, false, "Error loading site", fmt.Sprintf("%s/%s", name, site.Name()))
			} else {
				*size += uint64(len(tmpSite))
				*count++
				dir.files[site.Name()] = tmpSite
				logging.Debug(logging.SERVE, "Loaded site in cache", fmt.Sprintf("%s/%s", name, site.Name()))
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

func getSite(path string, host string) (*[]byte, int, error) {
	if util.GetConfig().Cache {
		if path == "/" {
			path = util.GetConfig().DefaultSite
		}
		for _, forbidden := range util.GetConfig().Forbidden.Endpoints {
			if strings.HasPrefix(path, forbidden+"/") || path == forbidden {
				site, code := GetErrorSite(Forbidden, host, path)
				return &site, code, errors.New(path + " Forbidden by Endpoints " + forbidden)
			}
		}

		for _, forbidden := range util.GetConfig().Forbidden.Regex {
			match, err := regexp.Match(forbidden, []byte(path))
			if err != nil {
				site, code := GetErrorSite(InternalServerError, host, path, fmt.Sprintf("Error checking forbidden regex"))
				return &site, code, err
			}
			if match {
				site, code := GetErrorSite(Forbidden, host, path)
				return &site, code, errors.New(path + " Forbidden by regex " + forbidden)
			}
		}

		pathSplit := strings.Split(path, "/")[1:]

		depth := len(pathSplit)

		dir := root
		for i := 0; i < depth-1; i++ {
			dir = dir.dirs[pathSplit[i]]
			if dir.files == nil {
				break
			}
		}
		site := dir.files[pathSplit[depth-1]]
		if site == nil {
			if _, ok := dir.dirs[pathSplit[depth-1]]; ok {
				site, code := GetErrorSite(NotFound, host, path, fmt.Sprintf("%s is no file, but a directory", pathSplit[depth-1]))
				return &site, code, errors.New(fmt.Sprintf("no site data for: %v", pathSplit))
			}
			site, code := GetErrorSite(NotFound, host, path)
			return &site, code, errors.New(fmt.Sprintf("no site data for: %s", pathSplit))
		}
		return &site, 200, nil

	} else {
		return nil, 500, errors.New("file reading not implemented")
	}
}

/*
CreateServe

Registers a handle for '/' to serve the DefaultSite
*/
func CreateServe() http.HandlerFunc {
	fun := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		msg, code, err := getSite(r.URL.Path, r.Host)

		if err != nil {
			logging.Err(logging.SERVE, err, false, fmt.Sprintf("Error getting site %s", r.URL.Path))
			w.WriteHeader(code)
		} else {
			fileSplit := strings.Split(r.URL.Path[1:], ".")
			filetype := fileSplit[len(fileSplit)-1]
			if val, exists := util.GetConfig().ContentTypes[filetype]; exists == true {
				w.Header().Set("Content-Type", val)
			}
		}
		searchTime := time.Now()
		_, er := w.Write(*msg)
		if er != nil {
			logging.Err(logging.SERVE, er, true, "Error writing response:")
		}
		go logging.LogAccess(code, int(time.Since(start).Microseconds()), int(searchTime.Sub(start).Microseconds()), err, er, r.TLS != nil, r.Method, r.URL.Path)
	}

	return fun
}
