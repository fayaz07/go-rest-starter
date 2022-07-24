package config

const (
	APP_ENV          = "APP_ENV"
	APP_PORT         = "APP_PORT"
	APP_DEFAULT_PORT = 7000

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
	ACCESS_TOKEN_EXPIRY_KEY  = "JWT_ACCESS_EXPIRY"
	REFRESH_TOKEN_EXPIRY_KEY = "JWT_REFRESH_EXPIRY"

	// environments
	ENV_KEY     = "env"
	PROD_ENV    = "prod"
	TEST_ENV    = "test"
	DEV_ENV     = "dev"
	STAGING_ENV = "staging"
)
