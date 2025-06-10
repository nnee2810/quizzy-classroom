package repository

import (
	"context"
	gorm2 "gorm.io/gorm"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
)

func (r *repositoryImpl) FilterJoinedClassrooms(ctx context.Context, params req.FilterJoinedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error) {
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
		Scopes(func(db *gorm2.DB) *gorm2.DB {
			if params.UserID != "" {
				db = db.Where("m.user_id = ?", params.UserID)
			}
			if params.Role != "" {
				db = db.Where("m.role = ?", params.Role)
			}
			return db
		}, gorm.Paginate(&pagination)).
		Preload("Members", "user_id = ?", params.UserID).
		Find(&classrooms).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = classrooms
	return &pagination, nil
}
