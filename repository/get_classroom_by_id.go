package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetClassroomByID(ctx context.Context, classroomID string) (*entity.ClassroomEntity, error) {
	var classroom entity.ClassroomEntity
	err := r.Provider.Db.WithContext(ctx).Where("id = ?", classroomID).First(&classroom).Error
	if err != nil {
		return nil, err
	}
	return &classroom, nil
}
