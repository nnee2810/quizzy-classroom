package usecase

import (
	"context"
	"quizzy-classroom/entity"
	_ "quizzy-classroom/provider"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type CreateClassroomUseCase interface {
	Execute(ctx context.Context, classroom *entity.ClassroomEntity) error
}

type createClassroomUseCaseImpl struct {
	Repo repository.Repository
}

func NewCreateClassroomUseCase(repo repository.Repository) CreateClassroomUseCase {
	return &createClassroomUseCaseImpl{
		Repo: repo,
	}
}

func (c *createClassroomUseCaseImpl) Execute(ctx context.Context, classroom *entity.ClassroomEntity) error {
	if err := c.Repo.CreateClassroom(ctx, classroom); err != nil {
		logger.Error("failed to create classroom", zap.Error(err))
		return err
	}
	return nil
}
