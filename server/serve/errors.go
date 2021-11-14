package serve

import (
	"fmt"
	"net/http"
	"strings"

	"Webgame/server/util"
)

type Errors uint16

const (
	Forbidden        Errors = http.StatusForbidden
	NotFound         Errors = http.StatusNotFound
	MethodNotAllowed Errors = http.StatusMethodNotAllowed // currently not in use
)

func GetErrorSite(error Errors, host string) ([]byte, int) {
	var replace = map[string]func() string{
		"%%code%%":   func() string { return fmt.Sprintf("%d|%s", error, http.StatusText(int(error))) },
		"%%public%%": func() string { return fmt.Sprintf("%s at Port %d", host, util.GetConfig().Port) },
	}

	var site string
	switch error {
	case Forbidden:
		site = getForbidden()
	case NotFound:
		site = getNotFound()
	case MethodNotAllowed:
		site = getMethodNotAllowed()
	default:
		site = getErrorNotFound()
	}
	replaceSite := getHeader() + site + getFoot()

	for repl, fun := range replace {
		replaceSite = strings.Replace(replaceSite, repl, fun(), -1)
	}
	return []byte(replaceSite), int(error)
}

func getForbidden() string {
	return "<h1>Forbidden</h1>" +
		"<p>You are not allowed to access this site.</p>"
}

func getNotFound() string {
	return "<h1>Not Found</h1>" +
		"<p>Site not found on server.</p>"
}

func getMethodNotAllowed() string {
	return "<h1>MethodNotAllowed</h1>" +
		"<p>Method not allowed.</p>"
}

func getErrorNotFound() string {
	return "<h1>Error not found</h1>" +
		"<p>Error not found</p>"
}

func getFoot() string {
	return "<hr>" +
		"<address>Golang Server %%public%%</address>"
}

func getHeader() string {
	return "<head> <title>%%code%%</title> </head>"
}
