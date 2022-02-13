package graph

import (
	"Server/logging"
	"Server/util"
)

/*
CheckAdmin

check if send hashcode exists or equals stored sha code
*/
func (r Resolver) CheckAdmin(code string) error {
	if util.GetConfig().Code == code {
		return nil
	}
	logging.Err(logging.API, nil, "An invalid code has been entered:", code)
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
