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

type FilterClassroomsUseCase interface {
	Execute(ctx context.Context, params req.FilterClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error)
}

type filterClassroomsUseCase struct {
	Repo repository.Repository
}

func NewFilterClassroomsUseCase(repo repository.Repository) FilterClassroomsUseCase {
	return &filterClassroomsUseCase{Repo: repo}
}

func (f *filterClassroomsUseCase) Execute(ctx context.Context, params req.FilterClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	result, err := f.Repo.FilterClassrooms(ctx, params)
	if err != nil {
		logger.Error("failed to filter classrooms", zap.Error(err))
		return nil, err
	}
	return result, nil
}
