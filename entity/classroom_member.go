package entity

import "github.com/nnee2810/mimi-core/record"

type ClassroomMemberRole string

const (
	ClassroomMemberRoleTeacher ClassroomMemberRole = "teacher"
	ClassroomMemberRoleStudent ClassroomMemberRole = "student"
)

type ClassroomMemberEntity struct {
	record.BaseEntity
	ClassroomID string              `json:"classroom_id" gorm:"uniqueIndex:idx_classroom_user"`
	UserID      string              `json:"user_id" gorm:"uniqueIndex:idx_classroom_user"`
	Role        ClassroomMemberRole `json:"role"`
	InvitedByID string              `json:"invited_by_id"`
	FullName    string              `json:"full_name" gorm:"-"`
}

func (*ClassroomMemberEntity) TableName() string {
	return "quizzy_classroom.classroom_members"
}
