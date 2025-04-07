package usecase

import (
	"context"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
	"quizzy-classroom/entity"
	_ "quizzy-classroom/provider"
	"quizzy-classroom/repository"
)

type CreateClassroomUseCase interface {
	Execute(ctx context.Context, classroom *entity.ClassroomEntity) error
}

type createClassroomUseCaseImpl struct {
	Repository repository.Repository
}

func NewCreateClassroomUseCase(repository repository.Repository) CreateClassroomUseCase {
	return &createClassroomUseCaseImpl{
		Repository: repository,
	}
}

func (c *createClassroomUseCaseImpl) Execute(ctx context.Context, classroom *entity.ClassroomEntity) error {
	if err := c.Repository.CreateClassroom(ctx, classroom); err != nil {
		logger.Error("failed to create classroom", zap.Error(err))
		return err
	}
	return nil
}
