package util

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type Type string

const (
	API    Type = "API"
	MAIN   Type = "MAIN"
	SERVE  Type = "SERVE"
	CONFIG Type = "CONFIG"
)

func Err(prefx Type, err error, printtrace bool, message ...interface{}) {
	log(prefx, "!", message...)
	if err != nil {
		log(prefx, "!", err.Error())
	}
	if printtrace {
		debug.PrintStack()
	}
}

func Log(prefx Type, message ...interface{}) {
	log(prefx, ">", message...)
}

func log(prefx Type, suffix string, message ...interface{}) {
	now := time.Now() // get this early.

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	var location string
	{
		file = file[strings.LastIndexByte(file, '/')+1:]
		var locationstretch = strconv.Itoa(int(GetConfig().Locationstretch))
		location = fmt.Sprintf("%-"+locationstretch+"s", fmt.Sprintf("%s:%d", file, line))
	}

	var prefix string
	if GetConfig().LogPrefix {
		var prefixstretch = strconv.Itoa(int(GetConfig().Prefixstretch))
		prefix = fmt.Sprintf("%-"+prefixstretch+"s", prefx)
	}

	var printstr string
	for _, mess := range message {
		printstr += fmt.Sprintf("%v", mess) + " "
	}

	os.Stdout.Write([]byte(fmt.Sprintf(
		"%s %s |%s %s %s \n",
		now.Format("2017-09-07 17:06:04.0000"),
		location, prefix, suffix, printstr,
	)))
}
