package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
)

func (r *repositoryImpl) FilterOwnedClassrooms(ctx context.Context, userID string, params req.FilterOwnedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	var pagination = record.Pagination[entity.ClassroomEntity]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, "created_at DESC"), // Mặc định sắp xếp theo ngày tạo mới nhất
	}

	var classrooms []entity.ClassroomEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Scopes(gorm.Paginate(&pagination)).
		Find(&classrooms).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = classrooms
	return &pagination, nil
}
