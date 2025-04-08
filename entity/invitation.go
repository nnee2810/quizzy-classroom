package entity

import (
	"github.com/nnee2810/mimi-core/record"
)

type InvitationEntity struct {
	record.BaseEntity
	ClassroomID string `json:"classroom_id"`
	SenderID    string `json:"sender_id" gorm:"index"`
	ReceiverID  string `json:"receiver_id" gorm:"index"`
	Status      string `json:"status" gorm:"index"`
}

func (InvitationEntity) TableName() string { return "quizzy_classroom.invitations" }
