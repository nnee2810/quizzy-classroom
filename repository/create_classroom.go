package repository

import (
	"context"
	"github.com/nnee2810/mimi-core/value"
	"gorm.io/gorm"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) CreateClassroom(ctx context.Context, classroom *entity.ClassroomEntity) error {
	return r.Provider.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(classroom).Error; err != nil {
			return err
		}

		if err := tx.Create(&entity.ClassroomMemberEntity{
			ClassroomID: value.GetValue(classroom.ID, ""),
			UserID:      classroom.UserID,
			Role:        entity.ClassroomMemberRoleTeacher,
		}).Error; err != nil {
			return err
		}

		return nil
	})
}
