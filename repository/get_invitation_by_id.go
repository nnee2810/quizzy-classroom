package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetInvitationByID(ctx context.Context, invitationID string) (*entity.InvitationEntity, error) {
	var invitation entity.InvitationEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Where("id = ?", invitationID).
		First(&invitation).
		Error; err != nil {
		return nil, err
	}
	return &invitation, nil
}
