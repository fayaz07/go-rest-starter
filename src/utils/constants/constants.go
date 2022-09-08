package constants

const (
	Status  = "status"
	Message = "message"
	Data    = "data"
	Error   = "error"

	Success = "success"
	Failed  = "failed"

	Invalid_Request_Body = "Invalid request body"
)

const (
	APP_ENV          = "APP_ENV"
	APP_PORT         = "APP_PORT"
	APP_DEFAULT_PORT = 7000

	// MongoDB
	DB_NAME     = "DB_NAME"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"

	// SendGrid
	SG_API_KEY  = "SENDGRID_API_KEY"
	SG_EMAIL    = "SENDGRID_EMAIL"
	SG_USERNAME = "SENDGRID_USERNAME"

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
