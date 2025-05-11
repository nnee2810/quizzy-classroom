package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) GetClassroomStat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	result, err := h.GetClassroomStatUseCase.Execute(c.Context(), userID)
	if err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, result)
}
