package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) UpdateInvitationStatus(ctx context.Context, invitationID string, status entity.InvitationStatus) error {
	return r.Provider.Db.
		WithContext(ctx).
		Model(&entity.InvitationEntity{}).
		Where("id = ?", invitationID).
		Update("status", status).
		Error
}
