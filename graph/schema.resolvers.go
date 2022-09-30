package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphqlGoKafka/graph/generated"
	"graphqlGoKafka/graph/model"
	"graphqlGoKafka/utils"
)

// CreatePreference is the resolver for the createPreference field.
func (r *mutationResolver) CreatePreference(ctx context.Context, input model.NewPreference) (string, error) {
	utils.SendPreferenceEvent(input)
	return "Ok", nil
}

// Preferences is the resolver for the preferences field.
func (r *queryResolver) Preferences(ctx context.Context) ([]*model.Preference, error) {
	var preferences []*model.Preference
	preferences = append(preferences, &model.Preference{ID: "5", Description: "pizza"})
	return preferences, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewPreference) (*model.Preference, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
