package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) CreateClassroom(ctx context.Context, classroom *entity.ClassroomEntity) error {
	return r.Provider.Db.WithContext(ctx).Create(classroom).Error
}
