package repository

import (
	"context"
	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
)

func (r *repositoryImpl) FilterClassrooms(ctx context.Context, params req.FilterClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	var pagination = record.Pagination[entity.ClassroomEntity]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, ""),
	}

	var classrooms []entity.ClassroomEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Scopes(gorm.Paginate(&pagination)).
		Find(&classrooms).
		Error; err != nil {
		return nil, err
	}

	userIDs := make([]string, len(classrooms))
	for i, classroom := range classrooms {
		userIDs[i] = classroom.UserID
	}

	userMap, err := r.GetUserProfilesByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	for i, member := range classrooms {
		classrooms[i].FullName = userMap[member.UserID].FullName
	}

	pagination.Rows = classrooms
	return &pagination, nil
}
