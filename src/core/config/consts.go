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
	SgAPIKey   = "SENDGRID_API_KEY"
	SgEmail    = "SENDGRID_EMAIL"
	SgUsername = "SENDGRID_USERNAME"

	// JWT
	accessTokenExpiry  = "JWT_ACCESS_EXPIRY"
	refreshTokenExpiry = "JWT_REFRESH_EXPIRY"

	// environments
	envConst    = "env"
	PROD_ENV    = "prod"
	TEST_ENV    = "test"
	DEV_ENV     = "dev"
	STAGING_ENV = "staging"
)
