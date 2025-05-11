package usecase

import (
	"context"
	"quizzy-classroom/entity"
	qerror "quizzy-classroom/error"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

type UpdateClassroomUseCase interface {
	Execute(ctx context.Context, classroomID string, userID string, classroom *entity.ClassroomEntity) error
}

type updateClassroomUseCaseImpl struct {
	Repo repository.Repository
}

func NewUpdateClassroomUseCase(repo repository.Repository) UpdateClassroomUseCase {
	return &updateClassroomUseCaseImpl{
		Repo: repo,
	}
}

func (u *updateClassroomUseCaseImpl) Execute(ctx context.Context, classroomID string, userID string, classroom *entity.ClassroomEntity) error {
	// Kiểm tra xem người dùng có phải là chủ sở hữu của lớp học không
	isOwner, err := u.Repo.IsClassroomOwner(ctx, classroomID, userID)
	if err != nil {
		logger.Error("failed to check classroom owner", zap.String("classroom id", classroomID), zap.String("user id", userID), zap.Error(err))
		return err
	}
	if !isOwner {
		return qerror.ErrNotClassroomOwner
	}

	// Cập nhật thông tin lớp học
	if err := u.Repo.UpdateClassroom(ctx, classroomID, classroom); err != nil {
		logger.Error("failed to update classroom", zap.String("classroom id", classroomID), zap.Error(err))
		return err
	}

	return nil
}
