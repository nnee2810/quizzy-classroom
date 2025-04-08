package repository

import (
	"context"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
	"quizzy-classroom/provider"

	"github.com/nnee2810/mimi-core/record"
)

type Repository interface {
	CreateClassroom(ctx context.Context, classroom *entity.ClassroomEntity) error
	FilterClassroomMembers(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMember], error)
}

type repositoryImpl struct {
	Provider *provider.Provider
}

func New(provider *provider.Provider) Repository {
	return &repositoryImpl{
		Provider: provider,
	}
}
