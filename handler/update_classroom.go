package handler

import (
	"errors"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/model/req"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (r *handlerImpl) UpdateClassroom(c *fiber.Ctx) error {
	classroomID := c.Params("classroom_id")
	userID := c.Locals("user_id").(string)

	var request req.UpdateClassroomReq
	if err := c.BodyParser(&request); err != nil {
		return res.BadRequest(c, err)
	}

	if err := r.UpdateClassroomUseCase.Execute(c.Context(), classroomID, userID, request); err != nil {
		if errors.Is(err, qerror.ErrNotClassroomOwner) {
			return res.Forbidden(c, err)
		}
		return res.InternalServerError(c, err)
	}

	return res.Success(c, nil)
}
