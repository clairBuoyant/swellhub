package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/clairBuoyant/swellhub/users/graph/model"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println("not implemented: Users - users")

	var users []*model.User
	users = append(users, &model.User{
		Name: "Weeble",
		ID:   "uuid_123",
	})
	return users, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
