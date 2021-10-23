package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type config struct {
	Port            uint16
	LogPrefix       bool
	Code            string
	Prefixstretch   int8
	Locationstretch int8
	Cache           bool
}

const (
	resourcesdir = "resources"
	Configfile   = resourcesdir + "/config.json"
	Sitesdir     = "site"
)

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() error {
	defaultConfig()

	data, err := ioutil.ReadFile(Configfile)
	if err != nil {
		Err(CONFIG, err, true, "Error reading", Configfile, "file")
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
	conf.Port = 0
	conf.LogPrefix = true
	conf.Code = fmt.Sprintf("this is supposed to be a secure code which should be overridden :Bonk: %d", rand.Int())
	conf.Prefixstretch = 0
	conf.Locationstretch = 0
	conf.Cache = true
}