package migrate

import (
	"quizzy-classroom/entity"
	"quizzy-classroom/model"
	"quizzy-classroom/provider"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func Run(serviceConfig *model.ServiceConfig) {
	provider, err := provider.Init(serviceConfig)
	if err != nil {
		logger.Error("failed to init provider", zap.Error(err))
		return
	}

	if err := provider.Db.AutoMigrate(
		&entity.ClassroomEntity{},
		&entity.ClassroomMemberEntity{},
		&entity.InvitationEntity{},
	); err != nil {
		logger.Error("failed to migrate database", zap.Error(err))
		return
	}

	logger.Info("migrate database success")
}
