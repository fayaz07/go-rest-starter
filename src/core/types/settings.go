package types

type AppSettings struct {
	AppName string `properties:"APP_NAME"`

	PublicDir        string `properties:"PUBLIC_DIR"`
	UploadsPublicDir string `properties:"UPLOADS_PUBLIC_DIR"`
	StaticPublicDir  string `properties:"STATIC_PUBLIC_DIR"`
	LogsDir          string `properties:"LOGS_DIR"`
	DotEnvsDir       string `properties:"DOT_ENVS_DIR"`
	DotEnvFile       string `properties:"DOT_ENV_FILE"`

	KeysDir            string `properties:"KEYS_DIR"`
	AccessTokenPvtKey  string `properties:"PVT_ACCESS_KEY"`
	AccessTokenPubKey  string `properties:"PUB_ACCESS_KEY"`
	RefreshTokenPvtKey string `properties:"PVT_REFRESH_KEY"`
	RefreshTokenPubKey string `properties:"PUB_REFRESH_KEY"`
}
