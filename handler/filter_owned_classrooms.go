package handler

import (
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) FilterOwnedClassrooms(c *fiber.Ctx) error {
	var params req.FilterOwnedClassroomsReq
	if err := c.QueryParser(&params); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&params); err != nil {
		return res.BadRequest(c, err)
	}

	// Lấy ID của người dùng hiện tại từ JWT token
	userID := c.Locals("user_id").(string)

	result, err := h.FilterOwnedClassroomsUseCase.Execute(c.Context(), userID, params)
	if err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, result)
}
