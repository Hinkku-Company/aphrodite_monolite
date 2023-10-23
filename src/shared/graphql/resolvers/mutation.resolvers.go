package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/graphql/generated"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/schema"
)

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, input schema.AccessToken) (bool, error) {
	return r.LoginModule.LogOut(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }