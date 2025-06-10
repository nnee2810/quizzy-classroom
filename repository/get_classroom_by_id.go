package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetClassroomByID(ctx context.Context, classroomID string, includeDetail bool) (*entity.ClassroomEntity, error) {
	var classroom entity.ClassroomEntity
	if err := r.Provider.Db.WithContext(ctx).Where("id = ?", classroomID).First(&classroom).Error; err != nil {
		return nil, err
	}

	if includeDetail {
		userMap, err := r.GetUserProfilesByIDs([]string{classroom.UserID})
		if err != nil {
			return nil, err
		}

		if user, ok := userMap[classroom.UserID]; ok {
			classroom.FullName = user.FullName
		}
	}

	return &classroom, nil
}
