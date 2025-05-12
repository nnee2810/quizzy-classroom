package entity

import (
	"github.com/nnee2810/mimi-core/record"
)

type InvitationStatus string

const (
	InvitationStatusPending  InvitationStatus = "pending"
	InvitationStatusAccepted InvitationStatus = "accepted"
	InvitationStatusRejected InvitationStatus = "rejected"
)

type InvitationEntity struct {
	record.BaseEntity
	ClassroomID string           `json:"classroom_id"`
	SenderID    string           `json:"sender_id" gorm:"index"`
	ReceiverID  string           `json:"receiver_id" gorm:"index"`
	Status      InvitationStatus `json:"status" gorm:"index"`

	Classroom *ClassroomEntity `json:"classroom" gorm:"foreignKey:ClassroomID"`
	FullName  string           `json:"full_name" gorm:"-"`
}

func (*InvitationEntity) TableName() string { return "quizzy_classroom.invitations" }
