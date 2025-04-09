package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) IsInvitationExisting(ctx context.Context, classroomID string, receiverID string) (bool, error) {
	var count int64
	if err := r.Provider.Db.
		WithContext(ctx).
		Model(&entity.InvitationEntity{}).
		Where("classroom_id = ? AND receiver_id = ? AND status IN ?", classroomID, receiverID, []entity.InvitationStatus{entity.InvitationStatusPending, entity.InvitationStatusAccepted}).
		Count(&count).
		Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
