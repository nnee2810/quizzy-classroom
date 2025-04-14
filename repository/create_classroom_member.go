package repository

import (
	"context"
	"quizzy-classroom/entity"

	"gorm.io/gorm/clause"
)

func (r *repositoryImpl) CreateClassroomMember(ctx context.Context, member *entity.ClassroomMember) error {
	// Sử dụng clause.OnConflict().DoNothing() để tránh lỗi khi thêm thành viên đã tồn tại
	return r.Provider.Db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(member).Error
}
