package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) IsClassroomOwner(ctx context.Context, classroomID string, userID string) (bool, error) {
	var classroom entity.ClassroomEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Where("id = ? AND user_id = ?", classroomID, userID).
		First(&classroom).
		Error; err != nil {
		// Nếu không tìm thấy bản ghi, có nghĩa là người dùng không phải là chủ sở hữu
		return false, nil
	}
	return true, nil
}
