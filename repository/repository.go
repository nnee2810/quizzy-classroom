package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/provider"
)

type Repository interface {
	CreateClassroom(ctx context.Context, classroom *entity.ClassroomEntity) error
}

type repositoryImpl struct {
	Provider *provider.Provider
}

func New(provider *provider.Provider) Repository {
	return &repositoryImpl{
		Provider: provider,
	}
}
