package util

import (
	"fmt"
	lg "log"
)

var stretch = "14"

func Log(prefx string, message ...interface{}) {
	prn := ""
	if GetConfig().LogPrefix {
		prn += fmt.Sprintf("%-"+stretch+"s|", prefx)
	}
	for _, mess := range message {
		prn += fmt.Sprintf("%v", mess)
	}
	lg.Println(prn)
}
