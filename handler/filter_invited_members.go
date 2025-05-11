package handler

import (
	"errors"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) FilterInvitedMembers(c *fiber.Ctx) error {
	// Lấy classroom_id từ URL
	classroomID := c.Params("classroom_id")
	if classroomID == "" {
		return res.BadRequest(c, errors.New("classroom_id is required"))
	}

	var query req.FilterInvitedMembersReq
	if err := c.QueryParser(&query); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&query); err != nil {
		return res.BadRequest(c, err)
	}

	result, err := h.FilterInvitedMembersUseCase.Execute(c.Context(), classroomID, c.Locals("user_id").(string), query)
	if err != nil {
		// Kiểm tra các loại lỗi
		if errors.Is(err, qerror.ErrNotClassroomOwner) {
			return res.BadRequest(c, errors.New("bạn không phải là chủ sở hữu của lớp học này"))
		}
		return res.InternalServerError(c, err)
	}

	return res.Success(c, result)
}
