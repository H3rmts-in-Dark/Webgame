package util

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type config struct {
	/*
		Port for the website must be between 0 and 65536
		this comes from the Dockerfile and should
		not get changed via the config file if used with Docker

		default:80
	*/
	PortHTTP uint16

	/*
		Port for the website must be between 0 and 65536
		this comes from the Dockerfile and should
		not get changed via the config file if used with Docker

		default:443
	*/
	PortHTTPS uint16

	/*
		Port used for the api must be between 0 and 65536
		should be different from Port to avoid trying to serve
		api by server

		default: 18266
	*/
	ApiPort uint16

	/*
		code needed to perform admin actions on the api

		default: random generated string
	*/
	Code string `yaml:"Code"`

	/*
		true: loads all files in the Sitesdir directory in cache

		false: loads file from Sitesdir directory every time it gets requested

				+ faster serve speed
				+ control over when users see new changes (show new files all at once)
				- changes to files must be loaded by sending a refresh request to the API
				- might load unnecessary files

		default: true
	*/
	Cache bool `yaml:"Cache"`

	/*
		enabled HTTP serving on this server on PortHTTP

		default: true
	*/
	EnableHTTP bool `yaml:"EnableHTTP"`

	/*
		enabled HTTPS with provided HTTPS certificate serving on this server on PortHTTPS

		default: true
	*/
	EnableHTTPS bool `yaml:"EnableHTTPS"`

	/*
		enabled HTTPS with provided HTTPS certificate serving on this server on ApiPort

		default: true
	*/
	ApiHTTPS bool `yaml:"ApiHTTPS"`

	/*
		change to serve root for serving files
		can be relative to the server main.go
		or absolute

		only this directory is served, but no underlying directory
		get served

		default: ./site
	*/
	SitesDir string `yaml:"SitesDir"`

	/*
		removes Debug logs from console if set to true
		can improve cache loading speed

		default: false
	*/
	Debug bool `yaml:"Debug"`

	/*
		adds a LogGroup to the log ( |CONFIG ) and adds a suffix to indicate
		the type of log (> for Normal, * for Debug, ! for Error) (<-- default)

		default: false
	*/
	LogPrefix bool `yaml:"LogPrefix"`

	/*
		stretches the prefix with LogGroup and > / * / ! to certain size

		can be ignored if LogPrefix is set to false

		default: 9 (fits the longest group)
	*/
	StretchPrefix uint8 `yaml:"StretchPrefix"`

	/*
		adds the file to the log where the Log method was called
		should be activated for debug purposes

		default: false
	*/
	LogFile bool `yaml:"LogFile"`

	/*
		stretches the filename:line number to certain size

		can be ignored if LogFile is set to false

		default: 16
	*/
	StretchFile uint8 `yaml:"StretchFile"`

	/*
		map of file extensions with the corresponding Content-Type

		{"css": "text/css; charset=utf-8} <-- example for .css files

		default: {}
	*/
	Headers map[string]string `yaml:"Headers"`

	/*
		which site to serve if no path was specified
		most likely be index.html

		default: "index.html"
	*/
	DefaultSite string `yaml:"DefaultSite"`

	/*
		list of files or paths which are not served return a Forbidden site

		{"api": true}  (<-- site/api returns a Forbidden site)

		default: []
	*/
	Forbidden []string `yaml:"Forbidden"`

	/*
		host of DB to connect to.
		Database to store logs, access logs, etc
	*/
	DBHost string `yaml:"DBHost"`

	/*
		user of DB to connect to.
		Database to store logs, access logs, etc
	*/
	DBUser string `yaml:"DBUser"`

	/*
		password of DBUser to connect to.
		Database to store logs, access logs, etc
	*/
	DBPassword string `yaml:"DBPassword"`

	/*
		database of DB to use.
		Database to store logs, access logs, etc
	*/
	DBDatabase string `yaml:"DBDatabase"`
}

const (
	ConfigFile = "server.yml"
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

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		Err(CONFIG, err, true, "Error unmarshalling configs")
		return err
	}

	// load some values from env
	if os.Getenv("PortHTTP") != "" {
		port, err := strconv.Atoi(os.Getenv("PortHTTP"))
		if err != nil {
			conf.PortHTTP = uint16(port)
		}
	}
	if os.Getenv("PortHTTPS") != "" {
		port, err := strconv.Atoi(os.Getenv("PortHTTPS"))
		if err != nil {
			conf.PortHTTPS = uint16(port)
		}
	}
	if os.Getenv("APIPORT") != "" {
		apiport, err := strconv.Atoi(os.Getenv("APIPORT"))
		if err != nil {
			conf.ApiPort = uint16(apiport)
		}
	}

	Log(CONFIG, "Loaded config:", fmt.Sprintf("%+v", conf))
	return nil
}

func defaultConfig() {
	conf.PortHTTP = 8080 // TODO revert
	conf.PortHTTPS = 8443
	conf.ApiPort = 18266
	conf.EnableHTTPS = true
	conf.EnableHTTP = true
	conf.ApiHTTPS = true
	conf.SitesDir = "./site"
	conf.Debug = false
	conf.LogFile = false
	conf.LogPrefix = false
	conf.Code = fmt.Sprintf("this is supposed to be a secure code which should be overridden :Bonk: %d", rand.Int())
	conf.StretchPrefix = 9
	conf.StretchFile = 16
	conf.Cache = true
	conf.Headers = map[string]string{}
	conf.DefaultSite = "index.html"
	conf.Forbidden = []string{}
	conf.DBHost = "no host provided"
	conf.DBUser = "no user provided"
	conf.DBPassword = "no password provided"
	conf.DBDatabase = "no database provided"
}
