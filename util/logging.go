package util

import (
	"fmt"
	lg "log"
)

var stretch = "12"

func Log(prefx string, message ...interface{}) {
	print := ""
	if GetConfig().LogPrefix {
		print += fmt.Sprintf("%-"+stretch+"s| ", prefx)
	}
	for _, mess := range message {
		print += fmt.Sprintf("%v", mess)
	}
	lg.Println(print)
}
