package api

import (
	"Webgame/serve"
	"Webgame/util"
)

/*
ProcessSiteReloadRequest

Process request to add Program to list of connections
*/
func ProcessSiteReloadRequest() (string, error) {
	return "success", serve.Loadsite()
}

/*
Checkadmin

check if send shacode exists or equals stored sha code
*/
func Checkadmin(js *map[string]interface{}) error {
	code, codeExists := (*js)["code"]
	if codeExists {
		if util.GetConfig().Code == code {
			return nil
		}
	}
	return &Permissionerror{}
}


/*
Permissionerror

Error thrown/returned when no admin priv are present
*/
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "no admin permissions"
}
