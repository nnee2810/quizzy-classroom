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

type FilterJoinedClassroomsUseCase interface {
	Execute(ctx context.Context, params req.FilterJoinedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error)
}

type filterJoinedClassroomsUseCaseImpl struct {
	Repo repository.Repository
}

func NewFilterJoinedClassroomsUseCase(repo repository.Repository) FilterJoinedClassroomsUseCase {
	return &filterJoinedClassroomsUseCaseImpl{
		Repo: repo,
	}
}

func (u *filterJoinedClassroomsUseCaseImpl) Execute(ctx context.Context, params req.FilterJoinedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	result, err := u.Repo.FilterJoinedClassrooms(ctx, params)
	if err != nil {
		logger.Error("failed to filter joined classrooms with role", zap.String("user id", params.UserID), zap.String("role", string(params.Role)), zap.Error(err))
		return nil, err
	}
	return result, nil
}
