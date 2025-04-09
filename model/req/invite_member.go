package req

type InviteMemberRequest struct {
	ClassroomID *string `json:"classroom_id" validate:"required"`
	ReceiverID  *string `json:"receiver_id" validate:"required"`
}
