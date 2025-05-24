package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetJoinedClassroomIDs(ctx context.Context, userID string) ([]string, error) {
	var classroomIDs []string

	if err := r.Provider.Db.
		WithContext(ctx).
		Model(&entity.ClassroomEntity{}).
		Where("user_id = ? AND role = ?", userID, entity.ClassroomMemberRoleStudent).
		Distinct().
		Pluck("classroom_id", &classroomIDs).
		Error; err != nil {
		return nil, err
	}

	return classroomIDs, nil
}
