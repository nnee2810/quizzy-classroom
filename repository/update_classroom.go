package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) UpdateClassroom(ctx context.Context, classroomID string, classroom *entity.ClassroomEntity) error {
	return r.Provider.Db.
		WithContext(ctx).
		Model(&entity.ClassroomEntity{}).
		Where("id = ?", classroomID).
		Updates(classroom).Error
}
