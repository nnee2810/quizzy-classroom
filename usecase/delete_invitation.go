package usecase

import (
	"context"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type DeleteInvitationUseCase interface {
	Execute(ctx context.Context, invitationID string, userID string) error
}

type deleteInvitationUseCaseImpl struct {
	Repo repository.Repository
}

func NewDeleteInvitationUseCase(repo repository.Repository) DeleteInvitationUseCase {
	return &deleteInvitationUseCaseImpl{
		Repo: repo,
	}
}

func (u *deleteInvitationUseCaseImpl) Execute(ctx context.Context, invitationID string, userID string) error {
	invitation, err := u.Repo.GetInvitationByID(ctx, invitationID)
	if err != nil {
		logger.Error("failed to get invitation", zap.String("invitation id", invitationID), zap.Error(err))
		return err
	}

	// Kiểm tra xem người dùng có phải là owner của classroom không
	isOwner, err := u.Repo.IsClassroomOwner(ctx, invitation.ClassroomID, userID)
	if err != nil {
		logger.Error("failed to check classroom owner", zap.String("classroom id", invitation.ClassroomID), zap.Error(err))
		return err
	}

	if !isOwner {
		return qerror.ErrNotClassroomOwner
	}

	if err = u.Repo.DeleteInvitation(ctx, invitationID); err != nil {
		logger.Error("failed to delete invitation", zap.String("invitation id", invitationID), zap.Error(err))
		return err
	}
	return nil
}
