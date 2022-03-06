package logging

import (
	"fmt"
	"time"

	"Server/util"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var cluster *gocql.ClusterConfig
var GQLSession gocqlx.Session
var session *gocql.Session

// DBInit Create and open DB Connection
func DBInit() {
	cluster = gocql.NewCluster(util.GetConfig().Database.Host)
	cluster.Port = int(util.GetConfig().Database.Port)
	cluster.Keyspace = util.GetConfig().Database.Database
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: util.GetConfig().Database.User,
		Password: util.GetConfig().Database.Password,
	}
	var err error
	GQLSession, err = gocqlx.WrapSession(cluster.CreateSession())
	session, err = cluster.CreateSession()
	if err != nil {
		Err(DB, err, "Error creating connection")
		panic(err)
	}
	Log(DB, "Connection established")
}

func LogAccess(code int, duration int, searchDuration int, error error, writeErr error, https bool, method string, uri string) {
	//language=SQL
	query := session.Query(
		"INSERT INTO server.access (id, uri, code, duration, searchDuration, method, https, error, writeErr) VALUES (?,?,?,?,?,?,?,?,?)",
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
		Err(DB, err, "Error inserting access into DB")
		Debug(query.Context())
	}
}

func LogAPIAccess(duration int, error error, request string) {
	//language=SQL
	query := session.Query(
		"INSERT INTO server.apiaccess (id, duration, error, request) VALUES (?,?,?,?)",
		gocql.TimeUUID(), duration, (func() interface{} {
			if error != nil {
				return error.Error()
			} else {
				return nil
			}
		})(), request)
	err := query.Exec()
	if err != nil {
		Err(DB, err, "Error inserting accessapi into DB")
		Debug(query.Context())
	}
}

func LoadMimeTypes() map[string]string {
	now := time.Now()
	//language=SQL
	iter := session.Query(
		"SELECT extension, mimetype FROM server.mime",
	).Iter()
	types := make(map[string]string, iter.NumRows())
	for {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		types[fmt.Sprintf("%s", row["extension"])] = fmt.Sprintf("%s", row["mimetype"])
	}
	if err := iter.Close(); err != nil {
		Err(DB, err, "Error loading Mime from DB")
		Debug(iter.Warnings())
	}
	Debug(DB, "Loaded Mime in", int(time.Since(now).Milliseconds()), "ms")
	return types
}
