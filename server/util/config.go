package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type config struct {
	/*
		Port for the website must be between 0 and 65536

		use something linke 80 or 443

		default:18265
	*/
	Port uint16

	/*
		Port used for the api must be between 0 and 65536
		should be different from Port to avoid trying to server
		api as a file

		default: 18266
	*/
	ApiPort uint16

	/*
		code needed to perform admin actions on the api
	*/
	Code string

	/*
		true: loads all files in the Sitesdir directory in cache

		false: loads file from Sitesdir directory every time it gets requested

				+ faster serve speed
				+ control over when users see new changes (show new files all at once)
				- changes to files must be loaded by sending a refresh request to the API
				- might load unnecessary files

		default: true
	*/
	Cache bool

	/*
		adds a LogGroup to the log ( |CONFIG ) and adds a suffix to indicate
		the type of log (> for Normal, * for Debug, ! for Error) (<-- default)

		default: false
	*/
	LogPrefix bool

	/*
		stretches the prefix with LogGroup and > / * / ! to certain size

		can be ignored if LogPrefix is set to false

		default: 9 (fits the longest group)
	*/
	StretchPrefix uint8

	/*
		adds the file to the log where the Log method was called
		should be activated for debug purposes

		default: false
	*/
	LogFile bool

	/*
		stretches the filename:line number to certain size

		can be ignored if LogFile is set to false

		default: 16
	*/
	StretchFile uint8

	/*
		map of file extensions with the corresponding Content-Type

		{"css": "text/css; charset=utf-8} <-- example for .css files

		default: {}
	*/
	Headers map[string]string

	/*
		which site to serve if no path was specified
		most likely be index.html

		default: "index.html"
	*/
	DefaultSite string

	/*
		list of files or paths which are not served return a Forbidden site

		{"api": true}  (<-- site/api returns a Forbidden site)

		default: []
	*/
	Forbidden []string
}

const (
	resourcesDir = "resources"
	ConfigFile   = resourcesDir + "/config.json"
	Sitesdir     = "site"
)

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() error {
	defaultConfig()

	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		Err(CONFIG, err, true, "Error reading", ConfigFile, "file")
		return err
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		Err(CONFIG, err, true, "Error unmarshalling configs")
		return err
	}
	Log(CONFIG, "Loaded config:", fmt.Sprintf("%+v", conf))
	return nil
}

func defaultConfig() {
	conf.Port = 18265
	conf.ApiPort = 18266
	conf.LogFile = false
	conf.LogPrefix = false
	conf.Code = fmt.Sprintf("this is supposed to be a secure code which should be overridden :Bonk: %d", rand.Int())
	conf.StretchPrefix = 9
	conf.StretchFile = 16
	conf.Cache = true
	conf.Headers = map[string]string{}
	conf.DefaultSite = "index.html"
	conf.Forbidden = []string{}
}
