package core

// DatabaseConfig - our database config
type DatabaseConfig struct {
	DbName     string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
}

// SendGridConfig - configuration of SendGrid
type SendGridConfig struct {
	SgAPIKey   string
	SgEmail    string
	SgUsername string
}

// JwtConfig - configuration of JWT keys
type JwtConfig struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

// AppConfig this is just a struct
type AppConfig struct {
	Initiated bool

	AppEnv  string
	AppPort int

	DB  *DatabaseConfig
	SG  *SendGridConfig
	Jwt *JwtConfig
}
