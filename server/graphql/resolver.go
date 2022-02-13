package graph

import (
	"context"
	"errors"

	"Server/graphql/generated"
	"Server/graphql/model"
	"Server/serve"
	"Server/util"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct{}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
