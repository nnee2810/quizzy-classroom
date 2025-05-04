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

	SupbaseUrl       string `mapstructure:"SUPABASE_URL"`
	SupabaseAnonKey  string `mapstructure:"SUPABASE_ANON_KEY"`
	SupbaseJwtSecret string `mapstructure:"SUPABASE_JWT_SECRET"`
}
