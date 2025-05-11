package handler

import (
	"errors"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/model/req"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) UpdateClassroom(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var request req.UpdateClassroomReq
	if err := c.BodyParser(&request); err != nil {
		return res.BadRequest(c, err)
	}

	if err := h.UpdateClassroomUseCase.Execute(c.Context(), request.ID, userID, &entity.ClassroomEntity{
		Name:        request.Name,
		Description: request.Description,
		AvatarUrl:   request.AvatarUrl,
	}); err != nil {
		if errors.Is(err, qerror.ErrNotClassroomOwner) {
			return res.Forbidden(c, err)
		}
		return res.InternalServerError(c, err)
	}

	return res.Success(c, nil)
}
