package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kunalpowar/gopggqlauth/server/graph/generated"
	"github.com/kunalpowar/gopggqlauth/server/graph/model"
	"github.com/kunalpowar/gopggqlauth/server/jwt"
	"github.com/kunalpowar/gopggqlauth/server/models/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	if err := user.Create(); err != nil {
		return "", fmt.Errorf("resolver: could not create user: %v", err)
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("resolver: could not generate token: %v", err)
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}

	correct, err := user.Authenticate()
	if err != nil {
		return "", fmt.Errorf("resolver: could not authenticate user: %v", err)
	}

	if !correct {
		return "", fmt.Errorf("resolver: username or password incorrect")
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
