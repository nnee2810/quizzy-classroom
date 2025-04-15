package req

import (
	"quizzy-classroom/entity"

	"github.com/nnee2810/mimi-core/model/req"
)

type FilterInvitedMembersReq struct {
	req.PaginationReq
	Status *entity.InvitationStatus `json:"status" query:"status"`
}
