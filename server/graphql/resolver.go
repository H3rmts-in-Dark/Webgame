package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"errors"
	"time"

	"Server/graphql/generated"
	"Server/graphql/model"
	"Server/logging"
	"Server/serve"
	"Server/util"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func (r *mutationResolver) ChangeSetting(ctx context.Context, setting model.NewSetting) (*model.Return, error) {
	panic("not implemented")
}

func (r *mutationResolver) ChangeAdmin(ctx context.Context, validation string, setting model.NewSetting) (*model.Return, error) {
	panic("not implemented")
}

func (r *mutationResolver) ReloadSites(ctx context.Context, validation string) (*model.Return, error) {
	err := r.CheckAdmin(validation)
	if err != nil {
		return &model.Return{Ok: false, Info: err.Error()}, err
	}
	if !util.GetConfig().Cache {
		return &model.Return{Ok: false, Info: err.Error()}, errors.New("caching deactivated")
	}
	err = serve.LoadSites()
	if err != nil {
		return &model.Return{Ok: false, Info: err.Error()}, err
	}
	return &model.Return{Ok: true, Info: "All Sites reloaded"}, nil
}

func (r *mutationResolver) ReloadSite(ctx context.Context, site string, validation string) (*model.Return, error) {
	panic("not implemented")
}

func (r *queryResolver) Settings(ctx context.Context) ([]*model.Setting, error) {
	panic("not implemented")
}

func (r *queryResolver) AdminSettings(ctx context.Context) ([]*model.Setting, error) {
	panic("not implemented")
}

func (r *queryResolver) Ping(ctx context.Context) (*model.Ping, error) {
	return &model.Ping{Uptime: int(time.Since(logging.Uptime).Seconds())}, nil
}

func (r *queryResolver) AccessLogs(ctx context.Context) ([]*model.Access, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.access").Iter()
	logs := make([]*model.Access, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.Access{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogs")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogs")
	return logs, nil
}

func (r *queryResolver) AccessLogsLimit(ctx context.Context, limit int) ([]*model.Access, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT id, code, duration, error, https, method, searchduration, uri, writeerr FROM server.access LIMIT ?", limit).Iter()
	logging.Log(logging.GRAPHQL, int(time.Since(now)))
	logs := make([]*model.Access, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.Access{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogsLimit")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsLimit")
	return logs, nil
}
func (r *queryResolver) AccessLogsByTime(ctx context.Context, from int, to int) ([]*model.Access, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.access WHERE id >= ? AND id <= ? ALLOW FILTERING",
		gocql.MinTimeUUID(time.Unix(int64(from), 0)),
		gocql.MaxTimeUUID(time.Unix(int64(to), 0)),
	).Iter()
	logs := make([]*model.Access, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.Access{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogsByTime")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsByTime")
	return logs, nil
}

func (r *queryResolver) AccessLogsByCode(ctx context.Context, from int, to int) ([]*model.Access, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.access WHERE code >= ? AND code <= ? ALLOW FILTERING", from, to).Iter()
	logs := make([]*model.Access, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.Access{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogsByCode")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsByCode")
	return logs, nil
}

func (r *queryResolver) AccessAPILogs(ctx context.Context) ([]*model.APIAccess, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.apiaccess").Iter()
	logs := make([]*model.APIAccess, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.APIAccess{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessAPILogs")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogs")
	return logs, nil
}

func (r *queryResolver) AccessAPILogsLimit(ctx context.Context, limit int) ([]*model.APIAccess, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.apiaccess LIMIT ?", limit).Iter()
	logs := make([]*model.APIAccess, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.APIAccess{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessAPILogsLimit")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogsLimit")
	return logs, nil
}

func (r *queryResolver) AccessAPILogsByTime(ctx context.Context, from int, to int) ([]*model.APIAccess, error) {
	now := time.Now()
	//language=SQL
	iter := Query(r.session, "SELECT * FROM server.apiaccess WHERE id >= ? AND id <= ? ALLOW FILTERING",
		gocql.MinTimeUUID(time.Unix(int64(from), 0)),
		gocql.MaxTimeUUID(time.Unix(int64(to), 0)),
	).Iter()
	logs := make([]*model.APIAccess, iter.Iter.NumRows())
	for i := 0; i < iter.Iter.NumRows(); i++ {
		log := model.APIAccess{}
		iter.StructScan(&log)
		logs[i] = &log
	}
	err := iter.Close()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error reading from DB")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessAPILogsByTime")
		return nil, errors.New("DB Error")
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogsByTime")
	return logs, nil
}

type Resolver struct {
	session *gocqlx.Session
}

type mutationResolver struct {
	*Resolver
}

type queryResolver struct {
	*Resolver
}

// // Mutation returns generated.MutationResolver implementation.
// func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func GenResolver() *Resolver {
	return &Resolver{
		session: &logging.GQLSession,
	}
}

func Query(session *gocqlx.Session, stmt string, params ...interface{}) *gocqlx.Queryx {
	return session.Query(stmt, nil).Bind(params...)
}
