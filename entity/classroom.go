package entity

import "github.com/nnee2810/mimi-core/record"

type ClassroomEntity struct {
	record.BaseEntity

	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func (c *ClassroomEntity) TableName() string {
	return "quizzy_classroom.classrooms"
}
