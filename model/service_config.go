package model

type ServiceConfig struct {
	Env  string `mapstructure:"ENV"`
	Port string `mapstructure:"PORT"`

	GrpcPort string `mapstructure:"GRPC_PORT"`

	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`

	SupabaseUrl            string `mapstructure:"SUPABASE_URL"`
	SupabaseServiceRoleKey string `mapstructure:"SUPABASE_SERVICE_ROLE_KEY"`
	SupabaseJwtSecret      string `mapstructure:"SUPABASE_JWT_SECRET"`
}
