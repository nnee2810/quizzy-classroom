package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) GetClassroomDetail(c *fiber.Ctx) error {
	classroomID := c.Params("classroom_id")
	if classroomID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "classroom_id is required")
	}

	classroom, err := h.GetClassroomUseCase.Execute(c.Context(), classroomID)
	if err != nil {
		return err
	}

	return res.Success(c, classroom)
}
