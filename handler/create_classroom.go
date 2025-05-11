package handler

import (
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) CreateClassroom(c *fiber.Ctx) error {
	var body req.CreateClassroomRequest

	if err := c.BodyParser(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := h.CreateClassroomUseCase.Execute(c.Context(), &entity.ClassroomEntity{
		UserID:      c.Locals("user_id").(string),
		Name:        body.Name,
		Description: body.Description,
		AvatarUrl:   body.AvatarUrl,
	}); err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, nil)
}
