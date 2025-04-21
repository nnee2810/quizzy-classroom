package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
)

func (r *repositoryImpl) FilterJoinedClassrooms(ctx context.Context, userID string, role entity.ClassroomMemberRole, params req.FilterJoinedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
	var pagination = record.Pagination[entity.ClassroomEntity]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, "created_at DESC"), // Mặc định sắp xếp theo ngày tạo mới nhất
	}

	var classrooms []entity.ClassroomEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Table("quizzy_classroom.classrooms AS c").
		Joins("JOIN quizzy_classroom.classroom_members AS m ON c.id = m.classroom_id").
		Where("m.user_id = ? AND m.role = ?", userID, role).
		Scopes(gorm.Paginate(&pagination)).
		Find(&classrooms).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = classrooms
	return &pagination, nil
}
