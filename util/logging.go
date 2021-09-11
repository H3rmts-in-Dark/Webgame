package util

import (
	"fmt"
	lg "log"
)

var stretch = "8"

func Log(prefx string, message ...interface{}) {
	var printstr string
	if GetConfig().LogPrefix {
		printstr += fmt.Sprintf("%-"+stretch+"s|", prefx)
	}
	for _, mess := range message {
		printstr += fmt.Sprintf("%v", mess) + " "
	}
	lg.Println(printstr)
}
