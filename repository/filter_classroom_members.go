package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"

	"github.com/nnee2810/mimi-core/gorm"
	"github.com/nnee2810/mimi-core/record"
	"github.com/nnee2810/mimi-core/value"
)

func (r *repositoryImpl) FilterClassroomMembers(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMemberEntity], error) {
	var pagination = record.Pagination[entity.ClassroomMemberEntity]{
		Page:  value.GetValue(params.Page, 0),
		Limit: value.GetValue(params.Limit, 0),
		Sort:  value.GetValue(params.Sort, ""),
	}

	var classroomMembers []entity.ClassroomMemberEntity
	if err := r.Provider.Db.
		WithContext(ctx).
		Where("classroom_id", classroomID).
		Scopes(gorm.Paginate(&pagination)).
		Find(&classroomMembers).
		Error; err != nil {
		return nil, err
	}

	userIDs := make([]string, len(classroomMembers))
	for i, member := range classroomMembers {
		userIDs[i] = member.UserID
	}

	userMap, err := r.GetUserProfilesByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	for i, member := range classroomMembers {
		classroomMembers[i].FullName = userMap[member.UserID].FullName
	}

	pagination.Rows = classroomMembers
	return &pagination, nil
}
