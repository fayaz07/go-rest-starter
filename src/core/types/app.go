package types

// AppConfig this is just a struct
type AppConfig struct {
	Initiated bool

	AppEnv  string
	AppPort int

	DB  *DatabaseConfig
	SG  *SendGridConfig
	Jwt *JwtConfig
}
