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

	// Fill owner name
	userIDs := make([]string, len(classrooms))
	for i, classroom := range classrooms {
		userIDs[i] = classroom.UserID
	}

	userMap, err := r.GetUserProfilesByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	for i, member := range classrooms {
		classrooms[i].FullName = userMap[member.UserID].FullName
	}

	// Fill member count
	classroomIDs := make([]string, len(classrooms))
	for i, classroom := range classrooms {
		classroomIDs[i] = *classroom.ID
	}

	var memberCounts []struct {
		ClassroomID string `json:"classroom_id"`
		Count       int64  `json:"count"`
	}

	if err := r.Provider.Db.WithContext(ctx).
		Model(&entity.ClassroomMemberEntity{}).
		Select("classroom_id, COUNT(*) as count").
		Where("classroom_id IN ?", classroomIDs).
		Group("classroom_id").
		Scan(&memberCounts).Error; err != nil {
		return nil, err
	}

	memberCountMap := make(map[string]int64)
	for _, v := range memberCounts {
		memberCountMap[v.ClassroomID] = v.Count
	}

	for i, classroomID := range classroomIDs {
		if count, ok := memberCountMap[classroomID]; ok {
			classrooms[i].MemberCount = count
		} else {
			classrooms[i].MemberCount = 0
		}
	}

	pagination.Rows = classrooms
	return &pagination, nil
}
