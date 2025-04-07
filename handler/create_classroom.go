package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/response"
	"github.com/nnee2810/mimi-core/value"
	"quizzy-classroom/entity"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"
)

func (r *handlerImpl) CreateClassroom(c *fiber.Ctx) error {
	var body req.CreateClassroomRequest

	if err := c.BodyParser(&body); err != nil {
		return response.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&body); err != nil {
		return response.BadRequest(c, err)
	}

	if err := r.CreateClassroomUseCase.Execute(c.Context(), &entity.ClassroomEntity{
		UserID:    value.Get(c.Locals("user_id").(string)),
		Name:      body.Name,
		AvatarUrl: body.AvatarUrl,
	}); err != nil {
		return response.InternalServerError(c, err)
	}

	return response.Success(c, nil)
}
