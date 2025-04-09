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

type FilterInvitationsUseCase interface {
	Execute(ctx context.Context, receiverID string, params req.FilterInvitationsReq) (*record.Pagination[entity.InvitationEntity], error)
}

type filterInvitationsUseCaseImpl struct {
	Repo repository.Repository
}

func NewFilterInvitationsUseCase(repo repository.Repository) FilterInvitationsUseCase {
	return &filterInvitationsUseCaseImpl{
		Repo: repo,
	}
}

func (u *filterInvitationsUseCaseImpl) Execute(ctx context.Context, receiverID string, params req.FilterInvitationsReq) (*record.Pagination[entity.InvitationEntity], error) {
	result, err := u.Repo.FilterInvitations(ctx, receiverID, params)
	if err != nil {
		logger.Error("failed to filter invitations", zap.String("receiver id", receiverID), zap.Error(err))
		return nil, err
	}
	return result, nil
}
