package repository

import (
	"context"
	"quizzy-classroom/provider"
)

type Repository interface {
	SignIn(ctx context.Context, id string) error
}

type repositoryImpl struct {
	Provider *provider.Provider
}

func New(provider *provider.Provider) Repository {
	return &repositoryImpl{
		Provider: provider,
	}
}
