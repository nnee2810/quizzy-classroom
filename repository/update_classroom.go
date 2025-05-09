package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
)

func (r *repositoryImpl) UpdateClassroom(ctx context.Context, classroomID string, req req.UpdateClassroomReq) error {
	return r.Provider.Db.
		WithContext(ctx).
		Model(&entity.ClassroomEntity{}).
		Where("id = ?", classroomID).
		Updates(map[string]interface{}{
			"name":        req.Name,
			"description": req.Description,
			"avatar_url":  req.AvatarUrl,
		}).Error
}
