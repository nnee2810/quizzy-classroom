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

type FilterClassroomMembersUseCase interface {
	Execute(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMember], error)
}

type filterClassroomMembersUseCase struct {
	Repo repository.Repository
}

func NewFilterClassroomMembersUseCase(repo repository.Repository) FilterClassroomMembersUseCase {
	return &filterClassroomMembersUseCase{Repo: repo}
}

func (f *filterClassroomMembersUseCase) Execute(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMember], error) {
	result, err := f.Repo.FilterClassroomMembers(ctx, classroomID, params)
	if err != nil {
		logger.Error("failed to filter classroom members", zap.String("classroom id", classroomID), zap.Error((err)))
		return nil, err
	}
	return result, nil
}
