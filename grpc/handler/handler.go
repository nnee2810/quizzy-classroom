package handler

import (
	"github.com/nnee2810/mimi-core/grpc/proto/gen"
	"quizzy-classroom/repository"
)

type Handler struct {
	gen.UnimplementedClassroomServer
	Repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{
		Repo: repo,
	}
}
