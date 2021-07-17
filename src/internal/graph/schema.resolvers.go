package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	generated1 "artics-api/src/internal/graph/generated"
	model1 "artics-api/src/internal/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model1.CreateUser) (*int, error) {
	return r.createUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model1.UpdateUser) (*model1.User, error) {
	return r.updateUser(ctx, input)
}

func (r *mutationResolver) SuspendUser(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Like(ctx context.Context, input model1.Like) (*int, error) {
	return r.like(ctx, input)
}

func (r *mutationResolver) Unlike(ctx context.Context, contentID int) (*int, error) {
	return r.unlike(ctx, contentID)
}

func (r *mutationResolver) Browse(ctx context.Context, contentID int) (*int, error) {
	return r.browse(ctx, contentID)
}

func (r *mutationResolver) CreateContent(ctx context.Context, input model1.CreateContent) (*int, error) {
	return r.createContent(ctx, input)
}

func (r *mutationResolver) UpdateContent(ctx context.Context, input model1.UpdateContent) (*model1.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContent(ctx context.Context, contentID int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Follow(ctx context.Context, followerID int) (*int, error) {
	return r.follow(ctx, followerID)
}

func (r *mutationResolver) Unfollow(ctx context.Context, followerID int) (*int, error) {
	return r.unfollow(ctx, followerID)
}

func (r *mutationResolver) AddToCloseUsers(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveFromCloseUsers(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddToMostFavorite(ctx context.Context, contentID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveFromMostFavorite(ctx context.Context, contentID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model1.User, error) {
	return r.currentUser(ctx)
}

func (r *queryResolver) Contents(ctx context.Context, id *int, kind model1.ContentKind) ([]*model1.Content, error) {
	return r.contents(ctx, id, kind)
}

func (r *queryResolver) Content(ctx context.Context, id int) (*model1.Content, error) {
	return r.content(ctx, id)
}

func (r *queryResolver) LikedBy(ctx context.Context, contentID int) ([]*model1.User, error) {
	return r.likedBy(ctx, contentID)
}

func (r *queryResolver) Notices(ctx context.Context) ([]model1.Notice, error) {
	return r.notices(ctx)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model1.User, error) {
	return r.user(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, kind model1.UserKind) ([]*model1.User, error) {
	return r.users(ctx, kind)
}

func (r *queryResolver) Favorites(ctx context.Context, id *int, kind model1.FavoriteKind) ([]*model1.Favorite, error) {
	return r.favorites(ctx, id, kind)
}

func (r *queryResolver) SearchUsers(ctx context.Context, words string) ([]*model1.User, error) {
	return r.searchUsers(ctx, words)
}

func (r *queryResolver) SearchContents(ctx context.Context, words string) ([]*model1.Content, error) {
	return r.searchContents(ctx, words)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) FavoriteBy(ctx context.Context, contentID int) ([]*model1.User, error) {
	panic(fmt.Errorf("not implemented"))
}
