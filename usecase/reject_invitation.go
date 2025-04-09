package usecase

import (
	"context"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type RejectInvitationUseCase interface {
	Execute(ctx context.Context, invitationID string, userID string) error
}

type rejectInvitationUseCaseImpl struct {
	Repo repository.Repository
}

func NewRejectInvitationUseCase(repo repository.Repository) RejectInvitationUseCase {
	return &rejectInvitationUseCaseImpl{
		Repo: repo,
	}
}

func (u *rejectInvitationUseCaseImpl) Execute(ctx context.Context, invitationID string, userID string) error {
	// Lấy thông tin invitation
	invitation, err := u.Repo.GetInvitationByID(ctx, invitationID)
	if err != nil {
		logger.Error("failed to get invitation", zap.String("invitation id", invitationID), zap.Error(err))
		return err
	}

	// Kiểm tra xem người dùng có phải là người nhận lời mời không
	if invitation.ReceiverID != userID {
		return qerror.ErrNotInvitationReceiver
	}

	// Kiểm tra xem invitation có đang ở trạng thái pending không
	if invitation.Status != entity.InvitationStatusPending {
		return qerror.ErrInvitationNotPending
	}

	// Cập nhật trạng thái invitation thành rejected
	if err := u.Repo.UpdateInvitationStatus(ctx, invitationID, entity.InvitationStatusRejected); err != nil {
		logger.Error("failed to update invitation status", zap.String("invitation id", invitationID), zap.Error(err))
		return err
	}

	return nil
}
