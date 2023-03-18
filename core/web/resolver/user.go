package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"goweb/core/types"
)

type UserResolver struct {
	user types.DBUser
}

func NewUser(user types.DBUser) *UserResolver {
	return &UserResolver{user: user}
}

// ID resolves the user's unique identifier.
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *UserResolver) Name() string {
	return r.user.Name
}

func (r *UserResolver) Age() int {
	return int(r.user.Age)
}

// CreatedAt resolves the chain's created at field.
func (r *UserResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.user.CreatedAt}
}

// UpdatedAt resolves the chain's updated at field.
func (r *UserResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.user.UpdatedAt}
}

type UserPayloadResolver struct {
	user types.DBUser
	NotFoundErrorUnionType
}

func NewUserPayload(user types.DBUser, err error) *UserPayloadResolver {
	e := NotFoundErrorUnionType{err: err, message: "chain not found", isExpectedErrorFn: nil}
	return &UserPayloadResolver{user: user, NotFoundErrorUnionType: e}
}

func (r *UserPayloadResolver) ToUser() (*UserResolver, bool) {
	if r.err != nil {
		return nil, false
	}

	return NewUser(r.user), true
}
