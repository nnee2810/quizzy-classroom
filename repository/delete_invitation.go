package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) DeleteInvitation(ctx context.Context, invitationID string) error {
	return r.Provider.Db.WithContext(ctx).
		Where("id = ?", invitationID).
		Delete(&entity.InvitationEntity{}).
		Error
}
