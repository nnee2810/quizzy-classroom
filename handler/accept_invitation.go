package handler

import (
	"errors"
	qerror "quizzy-classroom/error"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
	"gorm.io/gorm"
)

func (r *handlerImpl) AcceptInvitation(c *fiber.Ctx) error {
	// Lấy ID của invitation từ URL
	invitationID := c.Params("id")
	if invitationID == "" {
		return res.BadRequest(c, errors.New("id is required"))
	}

	// Lấy ID của người dùng hiện tại từ JWT token
	userID := c.Locals("user_id").(string)

	if err := r.AcceptInvitationUseCase.Execute(c.Context(), invitationID, userID); err != nil {
		// Kiểm tra các loại lỗi
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return res.BadRequest(c, errors.New("không tìm thấy lời mời"))
		case errors.Is(err, qerror.ErrNotInvitationReceiver):
			return res.BadRequest(c, errors.New("bạn không phải là người nhận lời mời này"))
		case errors.Is(err, qerror.ErrInvitationNotPending):
			return res.BadRequest(c, errors.New("lời mời không ở trạng thái chờ xử lý"))
		default:
			return res.InternalServerError(c, err)
		}
	}

	return res.Success(c, nil)
}
