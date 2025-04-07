package provider

import (
	"github.com/nedpals/supabase-go"
	"quizzy-classroom/model"

	"github.com/nnee2810/mimi-core/provider/database"
	"gorm.io/gorm"
)

type Provider struct {
	Db             *gorm.DB
	SupabaseClient *supabase.Client
}

func Init(serviceConfig *model.ServiceConfig) (*Provider, error) {
	db, err := database.NewPostgresDB(
		serviceConfig.DbHost,
		serviceConfig.DbPort,
		serviceConfig.DbUser,
		serviceConfig.DbPassword,
		serviceConfig.DbName,
	)
	if err != nil {
		return nil, err
	}

	supabaseClient := supabase.CreateClient(
		serviceConfig.SupbaseUrl,
		serviceConfig.SupabaseAnonKey,
	)

	return &Provider{
		Db:             db,
		SupabaseClient: supabaseClient,
	}, nil
}
