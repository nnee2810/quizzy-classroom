package req

type InviteMemberRequest struct {
	ClassroomID string `json:"classroom_id" validate:"required"`
	Email       string `json:"email" validate:"required"`
}
