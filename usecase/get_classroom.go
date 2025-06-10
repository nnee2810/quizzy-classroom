package usecase

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type GetClassroomUseCase interface {
	Execute(ctx context.Context, classroomID string) (*entity.ClassroomEntity, error)
}

type getClassroomUseCaseImpl struct {
	Repo repository.Repository
}

func NewGetClassroomUseCase(repo repository.Repository) GetClassroomUseCase {
	return &getClassroomUseCaseImpl{
		Repo: repo,
	}
}

func (g *getClassroomUseCaseImpl) Execute(ctx context.Context, classroomID string) (*entity.ClassroomEntity, error) {
	classroom, err := g.Repo.GetClassroomByID(ctx, classroomID, true)
	if err != nil {
		logger.Error("failed to get classroom by id", zap.Error(err), zap.String("classroom_id", classroomID))
		return nil, err
	}
	return classroom, nil
}
