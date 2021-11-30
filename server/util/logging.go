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

type logoptions struct {
	suffix    string
	colorcode string
}

var suffix = map[string]logoptions{"Debug": {"*", "\u001B[38;2;255;255;0m"}, "Normal": {">", "\u001B[38;2;0;255;0m"}, "Error": {"!", "\u001b[38;2;255;0;0m"}}

func Err(prefix LogGroup, err error, printTrace bool, message ...interface{}) {
	log(prefix, suffix["Error"], message...)
	if err != nil {
		log(prefix, suffix["Error"], err.Error())
	}
	if printTrace {
		debug.PrintStack()
	}
}

func Debug(message ...interface{}) {
	log("DEBUG", suffix["Debug"], message...)
}

func Log(prefix LogGroup, message ...interface{}) {
	log(prefix, suffix["Normal"], message...)
}

type LogWriter struct {
	Prefix LogGroup
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	Log(w.Prefix, string(p))
	return len(p), nil
}

func log(prefix LogGroup, logtype logoptions, message ...interface{}) {
	now := time.Now()

	var location string

	if GetConfig().LogFile {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}

		file = file[strings.LastIndexByte(file, '/')+1:] // convert to relative path
		var locationStretch = strconv.Itoa(int(GetConfig().StretchFile))
		location = fmt.Sprintf("%-"+locationStretch+"s", fmt.Sprintf("%s:%d", file, line))
	}

	var printPrefix string
	if GetConfig().LogPrefix {
		var prefixStretch = strconv.Itoa(int(GetConfig().StretchPrefix))
		printPrefix = fmt.Sprintf("%-"+prefixStretch+"s %s", prefix, logtype.suffix)
	}

	var printStr string
	for _, mess := range message {
		printStr += fmt.Sprintf("%v", mess) + " "
	}

	_, err := os.Stdout.Write([]byte(fmt.Sprintf(
		"%s%s %s|%s %s \u001b[0m\n",
		logtype.colorcode,
		now.Format("2006.01.02 15:04:05.00000"),
		location, printPrefix, printStr,
	)))
	if err != nil {
		return
	}
}
