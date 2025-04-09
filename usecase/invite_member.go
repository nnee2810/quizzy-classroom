package usecase

import (
	"context"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type InviteMemberUseCase interface {
	Execute(ctx context.Context, invitation *entity.InvitationEntity) error
}

type inviteMemberUseCaseImpl struct {
	Repo repository.Repository
}

func NewInviteMemberUseCase(repo repository.Repository) InviteMemberUseCase {
	return &inviteMemberUseCaseImpl{
		Repo: repo,
	}
}

func (u *inviteMemberUseCaseImpl) Execute(ctx context.Context, invitation *entity.InvitationEntity) error {
	// Kiểm tra xem người dùng có phải là chủ sở hữu của classroom không
	isOwner, err := u.Repo.IsClassroomOwner(ctx, invitation.ClassroomID, invitation.SenderID)
	if err != nil {
		logger.Error("failed to check classroom owner", zap.Error(err))
		return err
	}

	if !isOwner {
		return qerror.ErrNotClassroomOwner
	}

	// Kiểm tra xem người nhận đã có lời mời nào đang ở trạng thái pending hoặc accepted chưa
	exists, err := u.Repo.IsInvitationExisting(ctx, invitation.ClassroomID, invitation.ReceiverID)
	if err != nil {
		logger.Error("failed to check existing invitation", zap.Error(err))
		return err
	}

	if exists {
		return qerror.ErrReceiverAlreadyHasInvitation
	}

	if err := u.Repo.CreateInvitation(ctx, invitation); err != nil {
		logger.Error("failed to create invitation", zap.Error(err))
		return err
	}
	return nil
}
