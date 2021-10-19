package api

import (
	"Webgame/server/serve"
	"Webgame/server/util"
)

/*
ProcessSiteReloadRequest

Process request to add Program to list of connections
*/
func ProcessSiteReloadRequest() (string, error) {
	return "success", serve.Loadsites()
}

/*
CheckAdmin

check if send shacode exists or equals stored sha code
*/
func CheckAdmin(js *map[string]interface{}) error {
	code, codeExists := (*js)["code"]
	if codeExists {
		if util.GetConfig().Code == code {
			return nil
		}
		util.Err(util.API, nil, false, "An invalid code has been entered:", code)
	}
	return &Permissionerror{}
}

/*
Permissionerror

Error thrown/returned when no admin privileges are present
*/
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "No admin permissions"
}
