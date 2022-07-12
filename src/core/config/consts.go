package config

const (
	appEnv  = "APP_ENV"
	appPort = "APP_PORT"

	// MongoDB
	dBName     = "DB_NAME"
	dBHost     = "DB_HOST"
	dBPort     = "DB_PORT"
	dBUsername = "DB_USERNAME"
	dBPassword = "DB_PASSWORD"

	// SendGrid
	sgAPIKey   = "SENDGRID_API_KEY"
	sgEmail    = "SENDGRID_EMAIL"
	sgUsername = "SENDGRID_USERNAME"

	// JWT
	accessTokenExpiry  = "JWT_ACCESS_EXPIRY"
	refreshTokenExpiry = "JWT_REFRESH_EXPIRY"

	// environments
	prodEnv    = "prod"
	testEnv    = "test"
	devEnv     = "dev"
	stagingEnv = "staging"
)
