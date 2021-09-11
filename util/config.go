package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type config struct {
	Port      int32
	LogPrefix bool
	Code      string
}

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() {
	data, err := ioutil.ReadFile("util/config.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defaultConfig()
	
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	
	Log("Config", "loaded config: ", fmt.Sprintf("%+v", conf))
}

func defaultConfig() {
	conf.Port = -1
	conf.LogPrefix = true
}
