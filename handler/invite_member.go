package handler

import (
	"errors"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
	"github.com/nnee2810/mimi-core/value"
)

func (h *handlerImpl) InviteMember(c *fiber.Ctx) error {
	var body req.InviteMemberRequest

	if err := c.BodyParser(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := h.InviteMemberUseCase.Execute(c.Context(), &entity.InvitationEntity{
		ClassroomID: value.GetValue(body.ClassroomID, ""),
		SenderID:    c.Locals("user_id").(string),
		ReceiverID:  value.GetValue(body.ReceiverID, ""),
		Status:      entity.InvitationStatusPending,
	}); err != nil {
		// Kiểm tra các loại lỗi
		switch {
		case errors.Is(err, qerror.ErrReceiverAlreadyHasInvitation):
			return res.BadRequest(c, errors.New("người nhận đã có lời mời đang chờ xử lý hoặc đã được chấp nhận"))
		case errors.Is(err, qerror.ErrNotClassroomOwner):
			return res.BadRequest(c, errors.New("bạn không phải là chủ sở hữu của lớp học này"))
		default:
			return res.InternalServerError(c, err)
		}
	}

	return res.Success(c, nil)
}
