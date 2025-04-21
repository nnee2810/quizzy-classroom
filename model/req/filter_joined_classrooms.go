package req

import (
	"quizzy-classroom/entity"

	"github.com/nnee2810/mimi-core/model/req"
)

type FilterJoinedClassroomsReq struct {
	req.PaginationReq
	Role entity.ClassroomMemberRole `json:"role" query:"role" validate:"required"`
}
