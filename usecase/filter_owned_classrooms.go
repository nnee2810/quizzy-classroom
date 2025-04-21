package usecase

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"github.com/nnee2810/mimi-core/record"
	"go.uber.org/zap"
)

type FilterOwnedClassroomsUseCase interface {
	Execute(ctx context.Context, userID string, params req.FilterOwnedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error)
}

type filterOwnedClassroomsUseCaseImpl struct {
	Repo repository.Repository
}

func NewFilterOwnedClassroomsUseCase(repo repository.Repository) FilterOwnedClassroomsUseCase {
	return &filterOwnedClassroomsUseCaseImpl{
		Repo: repo,
	}
}

func (u *filterOwnedClassroomsUseCaseImpl) Execute(ctx context.Context, userID string, params req.FilterOwnedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	result, err := u.Repo.FilterOwnedClassrooms(ctx, userID, params)
	if err != nil {
		logger.Error("failed to filter owned classrooms", zap.String("user id", userID), zap.Error(err))
		return nil, err
	}
	return result, nil
}
