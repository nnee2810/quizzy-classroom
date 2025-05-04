package handler

import (
	"context"
	"github.com/nnee2810/mimi-core/grpc/proto/gen"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func (h *Handler) IsClassroomOwner(ctx context.Context, req *gen.IsClassroomOwnerRequest) (*gen.IsClassroomOwnerResponse, error) {
	ok, err := h.Repo.IsClassroomOwner(ctx, req.ClassroomID, req.UserID)
	if err != nil {
		logger.Error("failed to check classroom owner", zap.Error(err))
		return nil, err
	}
	return &gen.IsClassroomOwnerResponse{IsOwner: ok}, nil
}
