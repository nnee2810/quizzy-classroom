package handler

import (
	"context"

	"github.com/nnee2810/mimi-core/grpc/proto/gen"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func (h *Handler) GetJoinedClassroomIDs(ctx context.Context, req *gen.GetJoinedClassroomIDsRequest) (*gen.GetJoinedClassroomIDsResponse, error) {
	classroomIDs, err := h.Repo.GetJoinedClassroomIDs(ctx, req.UserID)
	if err != nil {
		logger.Error("failed to get joined classroom ids", zap.Error(err))
		return nil, err
	}
	return &gen.GetJoinedClassroomIDsResponse{ClassroomIDs: classroomIDs}, nil
}
