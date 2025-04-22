package entity

import "github.com/nnee2810/mimi-core/record"

type ClassroomMemberRole string

const (
	ClassroomMemberRoleTeacher ClassroomMemberRole = "teacher"
	ClassroomMemberRoleStudent ClassroomMemberRole = "student"
)

type ClassroomMember struct {
	record.BaseEntity
	ClassroomID string              `json:"classroom_id" gorm:"uniqueIndex:idx_classroom_user"`
	UserID      string              `json:"user_id" gorm:"uniqueIndex:idx_classroom_user"`
	Role        ClassroomMemberRole `json:"role"`
	InvitedByID string              `json:"invited_by_id"`
}

func (*ClassroomMember) TableName() string {
	return "quizzy_classroom.classroom_members"
}
