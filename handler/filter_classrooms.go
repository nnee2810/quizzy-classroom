package handler

import (
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) FilterClassrooms(c *fiber.Ctx) error {
	var params req.FilterClassroomsReq
	if err := c.QueryParser(&params); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&params); err != nil {
		return res.BadRequest(c, err)
	}

	result, err := h.FilterClassroomsUseCase.Execute(c.Context(), params)
	if err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, result)
}
