package logging

import (
	"Server/util"

	"github.com/gocql/gocql"
)

var session *gocql.Session

// DBInit Create and open DB Connection
func DBInit() {
	cluster := gocql.NewCluster(util.GetConfig().Database.Host)
	cluster.Port = int(util.GetConfig().Database.Port)
	cluster.Keyspace = util.GetConfig().Database.Database
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: util.GetConfig().Database.User,
		Password: util.GetConfig().Database.Password,
	}
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		Err(DB, err, true, "Error creating connection")
		panic(err)
	}
	Log(DB, "Connection established")
}

func LogAccess(code int, duration int, searchDuration int, error error, writeErr error, https bool, method string, uri string) {
	query := session.Query(
		"INSERT INTO access (id, uri, code, duration, searchDuration, method, https, error, writeErr) VALUES (?,?,?,?,?,?,?,?,?)",
		gocql.TimeUUID(), uri, code, duration, searchDuration, method, https, (func() interface{} {
			if error != nil {
				return error.Error()
			} else {
				return nil
			}
		})(), (func() interface{} {
			if writeErr != nil {
				return writeErr.Error()
			} else {
				return nil
			}
		})())
	err := query.Exec()
	if err != nil {
		Err(DB, err, false, "Error inserting access into DB")
		Debug(query.Context())
	}
}
