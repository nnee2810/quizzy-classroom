package entity

import "github.com/nnee2810/mimi-core/record"

type ClassroomEntity struct {
	record.BaseEntity

	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
	FullName    string `json:"full_name,omitempty"`   // Full name of the user who created the classroom
	MemberCount int64  `json:"member_count" gorm:"-"` // Number of members in the classroom
}

func (*ClassroomEntity) TableName() string {
	return "quizzy_classroom.classrooms"
}
