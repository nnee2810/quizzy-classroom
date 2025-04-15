package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	gorm2 "github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
	"gorm.io/gorm"
)

func (r *repositoryImpl) FilterInvitedMembers(ctx context.Context, classroomID string, params req.FilterInvitedMembersReq) (*record.Pagination[entity.InvitationEntity], error) {
	var (
		pagination = record.Pagination[entity.InvitationEntity]{
			Page:  value.GetValue(params.Page, 0),
			Limit: value.GetValue(params.Limit, 0),
			Sort:  value.GetValue(params.Sort, ""),
		}
		invitations []entity.InvitationEntity
	)

	if err := r.Provider.Db.
		WithContext(ctx).
		Where("classroom_id = ?", classroomID).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if params.Status != nil {
				return db.Where("status = ?", *params.Status)
			}
			return db
		}, gorm2.Paginate(&pagination)).
		Find(&invitations).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = invitations
	return &pagination, nil
}
