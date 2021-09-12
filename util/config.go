package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Port            uint16
	LogPrefix       bool
	Code            string
	Prefixstretch   int8
	Locationstretch int8
}

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() error {
	defaultConfig()
	
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		Err(CONFIG, err, true,"Error reading", configfile, "File")
		return err
	}
	
	err = json.Unmarshal(data, &conf)
	if err != nil {
		Err(CONFIG, err,true, "Error Unmarshalling Configs")
		return err
	}
	Log(CONFIG, "Loaded config: ", fmt.Sprintf("%+v", conf))
	return nil
}

func defaultConfig() {
	conf.Port = 0
	conf.LogPrefix = true
	conf.Code = "Bonk this should be overridden immediately"
	conf.Prefixstretch = 0
	conf.Locationstretch = 0
}

const (
	configfile = "util/config.json"
)
