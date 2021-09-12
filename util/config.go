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

const (
	configfile = "resources/config.json"
)

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() error {
	defaultConfig()

	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		Err(CONFIG, err, true, "Error reading", configfile, "file")
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
	conf.Code = "this_is_supposed_to_be_a_secure_code_which_should_be_overridden"
	conf.Prefixstretch = 0
	conf.Locationstretch = 0
}
