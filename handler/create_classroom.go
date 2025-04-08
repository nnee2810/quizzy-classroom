package handler

import (
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
	"github.com/nnee2810/mimi-core/value"
)

func (r *handlerImpl) CreateClassroom(c *fiber.Ctx) error {
	var body req.CreateClassroomRequest

	if err := c.BodyParser(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&body); err != nil {
		return res.BadRequest(c, err)
	}

	if err := r.CreateClassroomUseCase.Execute(c.Context(), &entity.ClassroomEntity{
		UserID:    c.Locals("user_id").(string),
		Name:      value.GetValue(body.Name, ""),
		AvatarUrl: value.GetValue(body.AvatarUrl, ""),
	}); err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, nil)
}
