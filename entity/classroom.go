package entity

import "github.com/nnee2810/mimi-core/record"

type ClassroomEntity struct {
	record.BaseEntity

	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}

func (*ClassroomEntity) TableName() string {
	return "quizzy_classroom.classrooms"
}
