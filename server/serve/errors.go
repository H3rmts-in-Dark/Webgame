package serve

import (
	"fmt"
	"net/http"
	"runtime"
)

type Errors uint16

const (
	Forbidden           Errors = http.StatusForbidden
	NotFound            Errors = http.StatusNotFound
	InternalServerError Errors = http.StatusInternalServerError
	MethodNotAllowed    Errors = http.StatusMethodNotAllowed // currently not in use
)

func GetErrorSite(error Errors, host string, path string, additional ...string) ([]byte, int) {
	var site string
	switch error {
	case Forbidden:
		site = getForbidden(path, additional...)
	case NotFound:
		site = getNotFound(path, additional...)
	case MethodNotAllowed:
		site = getMethodNotAllowed(path, additional...)
	case InternalServerError:
		site = getInternalServerError(path, additional...)
	default:
		site = getErrorNotFound(path, additional...)
	}
	replaceSite := getHeader(error, http.StatusText(int(error))) + site + getFoot(host, runtime.Version(), runtime.GOOS)

	return []byte(replaceSite), int(error)
}

func getForbidden(path string, additional ...string) string {
	return fmt.Sprintf(`
%s
<p>You are not allowed to access this site.</p>
<p>%s</p>
	`, getTop("Forbidden", path), additional)
}

func getNotFound(path string, additional ...string) string {
	return fmt.Sprintf(`
%s
<p>Site not found on server.</p>
<p>%s</p>
	`, getTop("Not Found", path), additional)
}

func getMethodNotAllowed(path string, additional ...string) string {
	return fmt.Sprintf(`
%s
<p>Method not allowed.</p>
<p>%s</p>
	`, getTop("MethodNotAllowed", path), additional)
}

func getInternalServerError(path string, additional ...string) string {
	return fmt.Sprintf(`
%s
<p>An error happend while processing your Request</p>
<p>%s</p>
	`, getTop("InternalServerError", path), additional)
}

func getErrorNotFound(path string, additional ...string) string {
	return fmt.Sprintf(`
%s
<p>Error not found</p>
<p>%s</p>
	`, getTop("Error not found", path), additional)
}

func getFoot(host string, goInfo string, sys string) string {
	return fmt.Sprintf(`
<hr>
	<address>Golang Server at %s running %s on %s</address>
	`, host, goInfo, sys)
}

func getHeader(error Errors, status string) string {
	return fmt.Sprintf(`
<head>
	<title>%d | %s</title>
</head>
	`, error, status)
}

func getTop(message string, file string) string {
	return fmt.Sprintf(`
<div style="display: flex;align-items: center;justify-content: space-between;">
	<h1 style="margin-block: 0.2em;">%s</h1>
	<p>Error accessing %s</p>
</div>
	`, message, file)
}
