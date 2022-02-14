package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"Server/graphql/generated"
	"Server/graphql/model"
	"Server/logging"
	"Server/serve"
	"Server/util"

	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
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
	panic("not implemented")
}

func (r *queryResolver) AccessLogs(ctx context.Context) ([]*model.Access, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM access",
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogs")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogs")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.Access, len(Map))
	for i, m := range Map {
		log := model.Access{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogs")
	return logs, nil
}

func (r *queryResolver) AccessLogsLimit(ctx context.Context, limit *int) ([]*model.Access, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM access LIMIT ?", limit,
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogsLimit")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.Access, len(Map))
	for i, m := range Map {
		log := model.Access{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsLimit")
	return logs, nil
}
func (r *queryResolver) AccessLogsByTime(ctx context.Context, from int, to int) ([]*model.Access, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM access WHERE id >= ? AND id <= ? ALLOW FILTERING",
		gocql.MinTimeUUID(time.Unix(int64(from), 0)),
		gocql.MaxTimeUUID(time.Unix(int64(to), 0)),
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogsByTime")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.Access, len(Map))
	for i, m := range Map {
		log := model.Access{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsByTime")
	return logs, nil
}

func (r *queryResolver) AccessLogsByCode(ctx context.Context, from int, to int) ([]*model.Access, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM access WHERE code >= ? AND code <= ? ALLOW FILTERING",
		from, to,
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogsByCode")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.Access, len(Map))
	for i, m := range Map {
		log := model.Access{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessLogsByCode")
	return logs, nil
}

func (r *queryResolver) AccessAPILogs(ctx context.Context) ([]*model.APIAccess, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM apiaccess",
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogs")
		go logging.LogAPIAccess(int(time.Since(now)), err, "AccessLogs")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.APIAccess, len(Map))
	for i, m := range Map {
		log := model.APIAccess{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogs")
	return logs, nil
}

func (r *queryResolver) AccessAPILogsLimit(ctx context.Context, limit *int) ([]*model.APIAccess, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM apiaccess LIMIT ?", limit,
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogsLimit")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.APIAccess, len(Map))
	for i, m := range Map {
		log := model.APIAccess{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogsLimit")
	return logs, nil
}

func (r *queryResolver) AccessAPILogsByTime(ctx context.Context, from int, to int) ([]*model.APIAccess, error) {
	now := time.Now()
	iter := r.session.Query(
		"SELECT * FROM apiaccess WHERE id >= ? AND id <= ? ALLOW FILTERING",
		gocql.MinTimeUUID(time.Unix(int64(from), 0)),
		gocql.MaxTimeUUID(time.Unix(int64(to), 0)),
	).Iter()
	Map, err := iter.SliceMap()
	if err != nil {
		logging.Err(logging.GRAPHQL, err, "Error getting AccessLogsByTime")
		return nil, errors.New("SQL Error")
	}
	logs := make([]*model.APIAccess, len(Map))
	for i, m := range Map {
		log := model.APIAccess{}
		err := r.createFromMap(m, &log)
		if err != nil {
			logging.Err(logging.GRAPHQL, err, fmt.Sprintf("Error creating AccessLog from map %v", m))
			continue
		}
		logs[i] = &log
	}
	go logging.LogAPIAccess(int(time.Since(now)), nil, "AccessAPILogsByTime")
	return logs, nil
}

type Resolver struct {
	session *gocql.Session
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

func (r Resolver) createFromMap(data map[string]interface{}, int interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: model.Decode,
		TagName:    "json",
		Result:     int,
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	if err != nil {
		return nil
	}
	return nil
}

func GenResolver() *Resolver {
	return &Resolver{
		session: logging.GQLsession,
	}
}
