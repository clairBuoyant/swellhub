package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/clairBuoyant/swellhub/feedback/graph/model"
)

// CreateFeedback is the resolver for the createFeedback field.
func (r *mutationResolver) CreateFeedback(ctx context.Context, input model.FeedbackInput) (*model.Feedback, error) {
	panic(fmt.Errorf("not implemented: CreateFeedback - createFeedback"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }