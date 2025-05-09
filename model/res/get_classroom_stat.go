package res

type GetClassRoomStatResponse struct {
	OwnedCount             int64 `json:"owned_count"`
	JoinedCount            int64 `json:"joined_count"`
	PendingInvitationCount int64 `json:"pending_invitation_count"`
}
