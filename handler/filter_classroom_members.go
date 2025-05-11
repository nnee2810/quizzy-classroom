package handler

import (
	"errors"
	"quizzy-classroom/model/req"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/model/res"
)

func (h *handlerImpl) FilterClassroomMembers(c *fiber.Ctx) error {
	classroomID := c.Params("classroom_id")
	if classroomID == "" {
		return res.BadRequest(c, errors.New("classroom_id is required"))
	}

	var params req.FilterClassroomMembersReq
	if err := c.QueryParser(&params); err != nil {
		return res.BadRequest(c, err)
	}

	if err := util.Validate.Struct(&params); err != nil {
		return res.BadRequest(c, err)
	}

	result, err := h.FilterClassroomMembersUseCase.Execute(c.Context(), classroomID, params)
	if err != nil {
		return res.InternalServerError(c, err)
	}

	return res.Success(c, result)
}
