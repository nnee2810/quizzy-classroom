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
	GetClassroomByID(ctx context.Context, classroomID string, includeDetail bool) (*entity.ClassroomEntity, error)
	FilterClassrooms(ctx context.Context, params req.FilterClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error)
	FilterClassroomMembers(ctx context.Context, classroomID string, params req.FilterClassroomMembersReq) (*record.Pagination[entity.ClassroomMemberEntity], error)
	CreateInvitation(ctx context.Context, invitation *entity.InvitationEntity) error
	IsInvitationExisting(ctx context.Context, classroomID string, receiverID string) (bool, error)
	IsClassroomOwner(ctx context.Context, classroomID string, userID string) (bool, error)
	FilterInvitations(ctx context.Context, receiverID string, params req.FilterInvitationsReq) (*record.Pagination[entity.InvitationEntity], error)
	GetInvitationByID(ctx context.Context, invitationID string) (*entity.InvitationEntity, error)
	UpdateInvitationStatus(ctx context.Context, invitationID string, status entity.InvitationStatus) error
	CreateClassroomMember(ctx context.Context, member *entity.ClassroomMemberEntity) error
	FilterClassroomInvitations(ctx context.Context, classroomID string, params req.FilterClassroomInvitationsReq) (*record.Pagination[entity.InvitationEntity], error)
	FilterJoinedClassrooms(ctx context.Context, params req.FilterJoinedClassroomsReq) (*record.Pagination[entity.ClassroomEntity], error)
	GetClassroomStat(ctx context.Context, userID string) (ownedCount, joinedCount, pendingCount int64, err error)
	UpdateClassroom(ctx context.Context, classroomID string, classroom *entity.ClassroomEntity) error
	GetJoinedClassroomIDs(ctx context.Context, userID string) ([]string, error)
	DeleteInvitation(ctx context.Context, invitationID string) error

	GetUserIDByEmail(ctx context.Context, email string) (string, error)
	GetUserProfilesByIDs(userIDs []string) (map[string]entity.UserProfileEntity, error)
}

type repositoryImpl struct {
	Provider *provider.Provider
}

func New(provider *provider.Provider) Repository {
	return &repositoryImpl{
		Provider: provider,
	}
}
