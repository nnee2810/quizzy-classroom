package repository

import (
	"context"
	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
)

func (r *repositoryImpl) FilterClassroomMembers(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMember], error) {
	var pagination = record.Pagination[entity.ClassroomMember]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, ""),
	}

	var classroomMembers []entity.ClassroomMember
	if err := r.Provider.Db.
		WithContext(ctx).
		Where("classroom_id", classroomID).
		Scopes(gorm.Paginate(&pagination)).
		Find(&classroomMembers).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = classroomMembers
	return &pagination, nil
}
