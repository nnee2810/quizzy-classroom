package usecase

import (
	"context"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/model/req"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"github.com/nnee2810/mimi-core/record"
	"go.uber.org/zap"
)

type FilterInvitedMembersUseCase interface {
	Execute(ctx context.Context, classroomID string, userID string, params req.FilterInvitedMembersReq) (*record.Pagination[entity.InvitationEntity], error)
}

type filterInvitedMembersUseCaseImpl struct {
	Repo repository.Repository
}

func NewFilterInvitedMembersUseCase(repo repository.Repository) FilterInvitedMembersUseCase {
	return &filterInvitedMembersUseCaseImpl{
		Repo: repo,
	}
}

func (u *filterInvitedMembersUseCaseImpl) Execute(ctx context.Context, classroomID string, userID string, params req.FilterInvitedMembersReq) (*record.Pagination[entity.InvitationEntity], error) {
	// Kiểm tra xem người dùng có phải là chủ sở hữu của classroom không
	isOwner, err := u.Repo.IsClassroomOwner(ctx, classroomID, userID)
	if err != nil {
		logger.Error("failed to check classroom owner", zap.Error(err))
		return nil, err
	}

	if !isOwner {
		return nil, qerror.ErrNotClassroomOwner
	}

	// Lấy danh sách thành viên đã mời
	result, err := u.Repo.FilterInvitedMembers(ctx, classroomID, params)
	if err != nil {
		logger.Error("failed to filter invited members", zap.String("classroom id", classroomID), zap.Error(err))
		return nil, err
	}

	return result, nil
}
