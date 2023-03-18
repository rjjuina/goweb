package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"goweb/core/types"
	"time"
)

type Resolver struct {
}

func (r *Resolver) User(ctx context.Context, args struct{ ID graphql.ID }) (*UserPayloadResolver, error) {
	user := types.DBUser{
		ID:        "1",
		Name:      "roy",
		Age:       31,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return NewUserPayload(user, nil), nil
}
