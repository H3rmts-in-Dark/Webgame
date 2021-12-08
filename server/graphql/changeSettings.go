package graph

import (
	"errors"

	"Webgame/server/serve"
	"Webgame/server/util"
)

/*
ProcessSiteReloadRequest

Process request to add Program to list of connections
*/
func ProcessSiteReloadRequest() error {
	if util.GetConfig().Cache {
		return serve.LoadSites()
	} else {
		return errors.New("caching deactivated")
	}
}

/*
CheckAdmin

check if send hashcode exists or equals stored sha code
*/
func (r Resolver) CheckAdmin(code string) error {
	if util.GetConfig().Code == code {
		return nil
	}
	util.Err(util.API, nil, false, "An invalid code has been entered:", code)
	return &PermissionError{}
}

/*
PermissionError

Error thrown/returned when no admin privileges are present
*/
type PermissionError struct{}

func (m *PermissionError) Error() string {
	return "No admin permissions"
}
