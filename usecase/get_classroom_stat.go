package usecase

import (
	"context"
	"quizzy-classroom/model/res"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type GetClassroomStatUseCase interface {
	Execute(ctx context.Context, userID string) (*res.GetClassRoomStatResponse, error)
}

type getClassroomStatUseCaseImpl struct {
	Repo repository.Repository
}

func NewGetClassroomStatUseCase(repo repository.Repository) GetClassroomStatUseCase {
	return &getClassroomStatUseCaseImpl{
		Repo: repo,
	}
}

func (u *getClassroomStatUseCaseImpl) Execute(ctx context.Context, userID string) (*res.GetClassRoomStatResponse, error) {
	// Lấy thống kê từ repository
	ownedCount, joinedCount, pendingCount, err := u.Repo.GetClassroomStat(ctx, userID)
	if err != nil {
		logger.Error("failed to get classroom stat", zap.String("user id", userID), zap.Error(err))
		return nil, err
	}

	return &res.GetClassRoomStatResponse{
		OwnedCount:             ownedCount,
		JoinedCount:            joinedCount,
		PendingInvitationCount: pendingCount,
	}, nil
}
