package handler

import (
	"errors"
	qerror "quizzy-classroom/error"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
	"gorm.io/gorm"
)

func (h *handlerImpl) DeleteInvitation(c *fiber.Ctx) error {
	invitationID := c.Params("id")
	if invitationID == "" {
		return res.BadRequest(c, errors.New("id is required"))
	}

	userID := c.Locals("user_id").(string)

	if err := h.DeleteInvitationUseCase.Execute(c.Context(), invitationID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res.NotFound(c, errors.New("không tìm thấy lời mời"))
		}
		if errors.Is(err, qerror.ErrNotClassroomOwner) {
			return res.BadRequest(c, errors.New("bạn không phải là chủ sở hữu của lớp học này"))
		}
		return res.InternalServerError(c, err)
	}

	return res.Success(c, nil)
}
