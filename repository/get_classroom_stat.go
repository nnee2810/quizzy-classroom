package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetClassroomStat(ctx context.Context, userID string) (ownedCount, joinedCount, pendingCount int64, err error) {
	if err = r.Provider.Db.
		WithContext(ctx).
		Model(&entity.ClassroomEntity{}).
		Where("user_id = ?", userID).
		Count(&ownedCount).
		Error; err != nil {
		return
	}

	// Đếm số lớp học mà người dùng tham gia (với vai trò học sinh)

	if err = r.Provider.Db.
		WithContext(ctx).
		Model(&entity.ClassroomMemberEntity{}).
		Where("user_id = ? AND role = ?", userID, entity.ClassroomMemberRoleStudent).
		Count(&joinedCount).
		Error; err != nil {
		return
	}

	// Đếm số lời mời đang chờ
	if err = r.Provider.Db.
		WithContext(ctx).
		Model(&entity.InvitationEntity{}).
		Where("receiver_id = ? AND status = ?", userID, entity.InvitationStatusPending).
		Count(&pendingCount).
		Error; err != nil {
		return
	}

	return
}
