package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) CreateInvitation(ctx context.Context, invitation *entity.InvitationEntity) error {
	return r.Provider.Db.WithContext(ctx).Create(invitation).Error
}