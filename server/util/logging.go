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

type LogGroup string

const (
	API       LogGroup = "API"
	MAIN      LogGroup = "MAIN"
	SERVE     LogGroup = "SERVE"
	CONFIG    LogGroup = "CONFIG"
	SERVER    LogGroup = "SERVER"
	SERVERAPI LogGroup = "SERVERAPI"
)

var suffix = map[string]string{"Debug": "*", "Normal": ">", "Error": "!"}

func Err(prefx LogGroup, err error, printTrace bool, message ...interface{}) {
	log(prefx, suffix["Error"], message...)
	if err != nil {
		log(prefx, suffix["Error"], err.Error())
	}
	if printTrace {
		debug.PrintStack()
	}
}

func Debug(message ...interface{}) {
	log("DEBUG", suffix["Debug"], message...)
}

func Log(prefx LogGroup, message ...interface{}) {
	log(prefx, suffix["Normal"], message...)
}

type LogWriter struct {
	Prefix LogGroup
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	Log(w.Prefix, string(p))
	return len(p), nil
}

func log(prefx LogGroup, suffix string, message ...interface{}) {
	now := time.Now()

	var location string

	if GetConfig().LogFile {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}

		file = file[strings.LastIndexByte(file, '/')+1:] // convert to relative path
		var locationstretch = strconv.Itoa(int(GetConfig().StretchFile))
		location = fmt.Sprintf("%-"+locationstretch+"s", fmt.Sprintf("%s:%d", file, line))
	}

	var prefix string
	if GetConfig().LogPrefix {
		var prefixstretch = strconv.Itoa(int(GetConfig().StretchPrefix))
		prefix = fmt.Sprintf("%-"+prefixstretch+"s %s", prefx, suffix)
	}

	var printstr string
	for _, mess := range message {
		printstr += fmt.Sprintf("%v", mess) + " "
	}

	_, err := os.Stdout.Write([]byte(fmt.Sprintf(
		"%s %s|%s %s \n",
		now.Format("2006.01.02 15:04:05.00000"),
		location, prefix, printstr,
	)))
	if err != nil {
		return
	}
}
