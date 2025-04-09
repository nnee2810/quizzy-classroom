package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
)

func (r *repositoryImpl) FilterInvitations(ctx context.Context, receiverID string, params req.FilterInvitationsReq) (*record.Pagination[entity.InvitationEntity], error) {
	var pagination = record.Pagination[entity.InvitationEntity]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, "created_at DESC"), // Mặc định sắp xếp theo ngày tạo mới nhất
	}

	query := r.Provider.Db.
		WithContext(ctx).
		Where("receiver_id = ?", receiverID)

	// Lọc theo status nếu có
	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}

	var invitations []entity.InvitationEntity
	if err := query.
		Scopes(gorm.Paginate(&pagination)).
		Find(&invitations).
		Error; err != nil {
		return nil, err
	}

	pagination.Rows = invitations
	return &pagination, nil
}
